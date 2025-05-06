package models

type Order struct {
	ID       string  `json:"id"`
	Item     string  `json:"item"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
