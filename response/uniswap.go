package response

type Price struct {
	Datedata string `json:"date_data"`
	Price    string `json:"price"`
}

func NewPrice(dateData string, price string) *Price {
	return &Price{
		Datedata: dateData,
		Price:    price,
	}
}
