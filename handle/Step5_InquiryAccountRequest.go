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

func InquiryAccount() {
	fmt.Println("=====> Step5: InquiryAccount transaction")

	timestamp := generateTimestamp()
	fmt.Println("Timestamp:", timestamp)

	requestData := common.AccountRequest{
		Merchant: common.Merchant{
			MerchantID:   merchantId,
			MerchantName: merchantName,
		},
		PaymentMethod:  "BRI",
		AccountNo:      "1234567890",
		HolderName:     "John Doe",
		AdditionalInfo: "",
	}

	jsonBytes, _ := json.Marshal(requestData)
	jsonStr := minify(string(jsonBytes))

	sha256Hash := sha256.Sum256([]byte(jsonStr))
	hashHex := hex.EncodeToString(sha256Hash[:])

	stringToSign := fmt.Sprintf("POST:%s:%s:%s:%s", inquiryAccountApi, accessTokenString, hashHex, timestamp)
	fmt.Println("String to sign:", stringToSign)
	signature, _ := hmacSHA512(stringToSign, merchantSecret)

	if err := postInquiryAccountRequest(jsonStr, timestamp, signature); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func postInquiryAccountRequest(jsonStr, timestamp, signature string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseURL, inquiryAccountApi), bytes.NewBuffer([]byte(jsonStr)))
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
