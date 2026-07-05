package main

import "fmt"

type PaymentMethod interface {
	pay(amount float64)
}
type Bkash struct {
	apiKey string
}

func (bk Bkash) pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Bkash", amount)
}

type PaymentService struct {
	method PaymentMethod
}

func (ps PaymentService) checkout() {
	ps.method.pay(100.00)
}
func main() {
	bkash := Bkash{apiKey: "4343hjkjj"}
	paymentService := PaymentService{
		method: bkash,
	}
	paymentService.checkout()
}
