package handle

import "time"

const (
	merchantId   = "10001"
	merchantName = "smilepay"

	accessTokenString = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJuYmYiOjE3NDA5ODAxNjAsImV4cCI6MTc0MDk4MTA2MCwiaWF0IjoxNzQwOTgwMTYwLCJNRVJDSEFOVF9JRCI6IjEwMDAxIn0.Yx5xu3ypU1I3Chp8n4qgSPLyPXGXOhIq7b1bwR1ivnE"
	merchantSecret    = "f4d768ef584ad56b5851ff071b5020f7a7e601ff912f7757cc3ef97e5808e44a"

	baseURL                 = "https://sandbox-gateway.smilepay.id"
	accessTokenApi          = "/v1.0/access-token/b2b"
	payInApi                = "/v1.0/transaction/pay-in"
	inquiryPaymentMethodApi = "/v1.0/disbursement/inquiry-paymentMethod"
	inquiryAccountApi       = "/v1.0/disbursement/inquiry-account"
	payOutApi               = "/v1.0/disbursement/cash-out"
)

func generateTimestamp() string {

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return ""
	}
	return time.Now().In(location).Format(time.RFC3339)
}
