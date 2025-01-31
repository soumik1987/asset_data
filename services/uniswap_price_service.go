package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/soumik1987/asset_price/helpers"
	"github.com/soumik1987/asset_price/models"
	"github.com/soumik1987/asset_price/requests"
	"github.com/soumik1987/asset_price/response"
)

// Need to move it to env file and take it from configurtion
// const uniswapAPI = "https://gateway.thegraph.com/api/apiKey/subgraphs/id/5zvR82QoaXYFyDEKLZ9t6v9adgnptxYpKpSbxtgVENFV"

const uniswapAPI = `https://gateway.thegraph.com/api/%s/subgraphs/id/5zvR82QoaXYFyDEKLZ9t6v9adgnptxYpKpSbxtgVENFV`

type IPriceService interface {
	FetchSpotPrices(tokenAddress string) ([]*response.Price, error)
}

type PriceService struct {
	client http.Client
	apiKey string
}

func NewPriceService(apiKey string) IPriceService {
	client := http.Client{}
	return &PriceService{
		client: client,
		apiKey: apiKey,
	}
}

// We have to move it to a separate Query file where all the queries will be loacted and configurable
const FetchPriceQuery = `
	{
		tokenDayDatas(
			first: 30, 
			orderBy: date, 
			orderDirection: desc, 
			where: { token: "%s" }
		) {
			date
			priceUSD
		}
	}
`

// We can make input as an array of string and fetch multiple
func (p *PriceService) FetchSpotPrices(tokenAddress string) ([]*response.Price, error) {
	query := requests.GraphQLQuery{
		Query: fmt.Sprintf(FetchPriceQuery, tokenAddress),
	}

	url := fmt.Sprintf(uniswapAPI, p.apiKey)
	fmt.Println(url, query)
	resp, err := helpers.ProcessRequest(query, p.client, url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var priceData models.PriceData
	if err := json.NewDecoder(resp.Body).Decode(&priceData); err != nil {
		return nil, err
	}

	var results []*response.Price
	for _, entry := range priceData.Data.TokenDayDatas {
		results = append(results, response.NewPrice(time.Unix(entry.Date, 0).Format("2006-01-02"), entry.PriceUSD))
	}
	fmt.Println("Success")
	return results, nil
}
