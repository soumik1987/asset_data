package response

type TokenDayDatas struct {
	Date     int64  `json:"date"`
	PriceUSD string `json:"priceUSD"`
}

type tokenData struct {
	TokenDayDatas []TokenDayDatas `json:"tokenDayDatas"`
}

type PriceData struct {
	Data tokenData `json:"data"`
}

type Price struct {
	Timestamp string `json:"timestamp"`
	Price    string `json:"price"`
}

func NewPrice(timestamp string, price string) *Price {
	return &Price{
		Timestamp: timestamp,
		Price:    price,
	}
}
