package common

type PayInRequest struct {
	OrderNo         string           `json:"orderNo"`
	Purpose         string           `json:"purpose"`
	ProductDetail   string           `json:"productDetail"`
	AdditionalParam string           `json:"additionalParam"`
	ItemDetailList  []ItemDetailList `json:"itemDetailList"`
	Money           Money            `json:"money"`
	Merchant        Merchant         `json:"merchant"`
	Payer           Payer            `json:"payer"`
	Receiver        Receiver         `json:"receiver"`
	BillingAddress  Address          `json:"billingAddress"`
	ShippingAddress Address          `json:"shippingAddress"`
	PaymentMethod   string           `json:"paymentMethod"`
	CallbackURL     string           `json:"callbackUrl,omitempty"`
	RedirectURL     string           `json:"redirectUrl,omitempty"`
}
