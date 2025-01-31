package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/soumik1987/asset_price/requests"
)

func ProcessRequest(query requests.GraphQLQuery, client http.Client, url string) (*http.Response, error) {
	requestBody, _ := json.Marshal(query)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
