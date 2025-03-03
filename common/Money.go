package common

type Money struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}
