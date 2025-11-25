package main

import (
	"fmt"
	"proj/payments"
	"proj/payments/methods"
)

func main() {

	method := methods.NewCrypto()

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("qwerty", 5)

	allinfo := paymentModule.AllInfo()

	fmt.Println(allinfo)

}
