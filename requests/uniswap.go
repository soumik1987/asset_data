package requests

// GraphQL query to get last 30 days of spot prices
type GraphQLQuery struct {
	Query string `json:"query"`
}

type PriceRequest struct {
	TokenAddress string `json:"token_address" query:"token_address"`
}
