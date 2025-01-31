package models

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
