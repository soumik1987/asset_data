package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/labstack/gommon/log"

	"github.com/soumik1987/asset_price/helpers"
	"github.com/soumik1987/asset_price/models"
	"github.com/soumik1987/asset_price/requests"
	"github.com/soumik1987/asset_price/response"
)

// Need to move it to env file and take it from configurtion
const uniswapAPI = `https://gateway.thegraph.com/api/%s/subgraphs/id/5zvR82QoaXYFyDEKLZ9t6v9adgnptxYpKpSbxtgVENFV`

type UniswapService struct {
	client      http.Client
	graphApiKey string
}

func NewUniswapService(graphApiKey string) models.IUniswapService {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	return &UniswapService{
		client:      client,
		graphApiKey: graphApiKey,
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
func (p UniswapService) FetchSpotPrices(tokenAddress string) ([]*response.Price, error) {
	query := requests.GraphQLQuery{
		Query: fmt.Sprintf(FetchPriceQuery, tokenAddress),
	}

	url := fmt.Sprintf(uniswapAPI, p.graphApiKey)
	resp, err := helpers.ProcessRequest(query, p.client, url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var priceData response.PriceData
	if err := json.NewDecoder(resp.Body).Decode(&priceData); err != nil {
		return nil, err
	}

	var results []*response.Price
	for _, entry := range priceData.Data.TokenDayDatas {
		results = append(results, response.NewPrice(time.Unix(entry.Date, 0).Format("2006-01-02"), entry.PriceUSD))
	}
	log.Info("API response completed")
	return results, nil
}
