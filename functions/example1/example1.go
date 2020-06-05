package example1

import "encoding/json"

func CalculateVat(totalPrice, percen float64) float64 {
	return totalPrice * (percen / 100)
}

type Product struct {
	ID    int     `json:"id"`
	Nmae  string  `json:"name"`
	Price float64 `json:"price"`
}

func JsonProduct(p Product) (string, error) {
	json, err := json.Marshal(p)
	return string(json), err
}
