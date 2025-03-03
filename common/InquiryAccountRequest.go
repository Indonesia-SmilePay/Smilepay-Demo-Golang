package common

type AccountRequest struct {
	Merchant       Merchant `json:"merchant"`
	PaymentMethod  string   `json:"paymentMethod"`
	AccountNo      string   `json:"accountNo"`
	HolderName     string   `json:"holderName"`
	AdditionalInfo string   `json:"additionalInfo"`
}
