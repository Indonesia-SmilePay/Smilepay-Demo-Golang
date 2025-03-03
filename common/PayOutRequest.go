package common

type TradePayoutReq struct {
	OrderNo         string           `json:"orderNo"`
	Purpose         string           `json:"purpose"`
	ProductDetail   string           `json:"productDetail"`
	AdditionalParam string           `json:"additionalParam"`
	ItemDetailList  []ItemDetailList `json:"itemDetailList"`
	BillingAddress  Address          `json:"billingAddress"`
	ShippingAddress Address          `json:"shippingAddress"`
	Money           Money            `json:"money"`
	Merchant        Merchant         `json:"merchant"`
	CallbackUrl     string           `json:"callbackUrl,omitempty"`
	RedirectUrl     string           `json:"redirectUrl,omitempty"`
	PaymentMethod   string           `json:"paymentMethod"`
	CashAccount     string           `json:"cashAccount"`
	Payer           Payer            `json:"payer"`
	Receiver        Receiver         `json:"receiver"`
}
