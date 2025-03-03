package main

import "SmilePay-Demo-Golang/handle"

func main() {

	//Step one, generate RSA key
	handle.GenerateRSA()

	//Step two, generate token
	handle.AccessToken()

	//Step three, call pay interface
	handle.PayIn()

	//step four, call inquiry interface
	handle.InquiryPaymentMethod()

	//step five, call inquiry account interface
	handle.InquiryAccount()

	//step six, call payout interface
	handle.PayOut()
}
