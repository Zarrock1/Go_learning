package methods

import (
	"fmt"
	"math/rand"
)

type Paypal struct{}

func NewPaypal() Paypal {

	return Paypal{}
}

func (p Paypal) Pay(usd int) int {
	fmt.Println("Оплата Paypal")
	fmt.Println("Размер оплаты:", usd, "USD")

	id := rand.Int()

	return id
}

func (p Paypal) Cancel(id int) {
	fmt.Println("Отмена операции Paypal")
}
