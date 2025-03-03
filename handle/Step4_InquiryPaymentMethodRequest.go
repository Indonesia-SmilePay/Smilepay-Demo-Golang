package handle

import (
	"SmilePay-Demo-Golang/common"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func InquiryPaymentMethod() {
	fmt.Println("=====> Step4: InquiryPaymentMethodRequest transaction")

	timestamp := generateTimestamp()
	fmt.Println("Timestamp:", timestamp)

	//generate inquiryPaymentMethodReq data
	requestData := common.InquiryPaymentMethodReq{
		Merchant: common.Merchant{
			MerchantID:   merchantId,
			MerchantName: merchantName,
		},
		AdditionalInfo: "",
	}

	jsonBytes, _ := json.Marshal(requestData)
	jsonStr := minify(string(jsonBytes))

	sha256Hash := sha256.Sum256([]byte(jsonStr))
	hashHex := hex.EncodeToString(sha256Hash[:])

	stringToSign := fmt.Sprintf("POST:%s:%s:%s:%s", inquiryPaymentMethodApi, accessTokenString, hashHex, timestamp)
	fmt.Println("String to sign:", stringToSign)
	signature, _ := hmacSHA512(stringToSign, merchantSecret)

	if err := postInquiryPaymentMethodRequest(jsonStr, timestamp, signature); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func postInquiryPaymentMethodRequest(jsonStr, timestamp, signature string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseURL, inquiryPaymentMethodApi), bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessTokenString)
	req.Header.Set("X-TIMESTAMP", timestamp)
	req.Header.Set("X-SIGNATURE", signature)
	req.Header.Set("X-PARTNER-ID", merchantId)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	log.Printf("Response: %s", responseBody)
	return nil
}
