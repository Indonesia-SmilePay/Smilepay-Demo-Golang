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
	"math/rand"
	"net/http"
	"time"
)

func PayIn() {
	fmt.Println("=====> Step3: PayIn transaction")

	timestamp := generateTimestamp()
	fmt.Println("Timestamp:", timestamp)

	rand.Seed(time.Now().UnixNano())
	orderNo := fmt.Sprintf("%s%d", merchantId, rand.Int63())

	requestData := common.PayInRequest{
		OrderNo:         orderNo,
		Purpose:         "Purpose for Transaction from Go SDK",
		ProductDetail:   "Product details",
		AdditionalParam: "Other descriptions",
		ItemDetailList: []common.ItemDetailList{
			{Name: "Mac A1", Quantity: 1, Price: 10000.00},
		},
		Money: common.Money{
			Currency: "IDR",
			Amount:   10000.00,
		},
		Merchant: common.Merchant{
			MerchantID:   merchantId,
			MerchantName: merchantName,
		},
		Payer: common.Payer{
			Name:    "test",
			Phone:   "0837984192",
			Address: "Jalan Pantai Mutiara TG6, Pluit, Jakarta",
			Email:   "integration@smilepay.id",
		},
		Receiver: common.Receiver{
			Name:    "smilepay",
			Phone:   "0837984192",
			Address: "Jl. Pluit Karang Ayu 1 No.B1 Pluit",
			Email:   "integration@smilepay.id",
		},
		BillingAddress: common.Address{
			CountryCode: "Indonesia",
			City:        "Jakarta",
			Address:     "Jl. Pluit Karang Ayu 1 No.B1 Pluit",
			Phone:       "0837984192",
			PostalCode:  "14450",
		},
		ShippingAddress: common.Address{
			CountryCode: "Indonesia",
			City:        "Jakarta",
			Address:     "Jl. Pluit Karang Ayu 1 No.B1 Pluit",
			Phone:       "0837984192",
			PostalCode:  "14450",
		},
		PaymentMethod: "BRI",
	}

	jsonBytes, _ := json.Marshal(requestData)
	jsonStr := minify(string(jsonBytes))

	sha256Hash := sha256.Sum256([]byte(jsonStr))
	hashHex := hex.EncodeToString(sha256Hash[:])

	stringToSign := fmt.Sprintf("POST:%s:%s:%s:%s", payInApi, accessTokenString, hashHex, timestamp)
	fmt.Println("String to sign:", stringToSign)
	signature, _ := hmacSHA512(stringToSign, merchantSecret)

	if err := postPayInRequest(jsonStr, timestamp, signature); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func postPayInRequest(jsonStr, timestamp, signature string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", baseURL, payInApi), bytes.NewBuffer([]byte(jsonStr)))
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
