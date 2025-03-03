package common

type Address struct {
	CountryCode string `json:"countryCode"`
	City        string `json:"city"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	PostalCode  string `json:"postalCode"`
}
