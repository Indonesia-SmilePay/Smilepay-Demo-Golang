package common

type InquiryPaymentMethodReq struct {
	Merchant       Merchant `json:"merchant"`
	AdditionalInfo string   `json:"additionalInfo"`
}
