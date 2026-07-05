package main

import "fmt"

type PaymentMethod interface {
	pay(amount float64)
}
type Bkash struct {
	apiKey string
}
type Nagad struct {
	apiKey string
}

func (bk *Bkash) pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Bkash", amount)
}
func (ng *Nagad) pay(amount float64) {
	fmt.Printf("Paying %.2f tk with Nagad", amount)
}

type PaymentService struct {
	method PaymentMethod
}

func NewPaymentService(method PaymentMethod) *PaymentService {
	return &PaymentService{
		method: method,
	}

}

func NewNagad(apiKey string) *Nagad {
	return &Nagad{
		apiKey: apiKey,
	}
}

func (ps PaymentService) checkout() {
	ps.method.pay(100.00)
}

type MockPaymentMethod struct {
}

func (mockPm MockPaymentMethod) pay(amount float64) {
	fmt.Println("Testing payment method")
}
func main() {
	// bkash := Bkash{apiKey: "4343hjkjj"}
	// paymentService1 := NewPaymentService(&bkash)
	// paymentService1.checkout()

	// nagad := NewNagad("ksddff83in")
	// paymentService2 := NewPaymentService(nagad)
	// paymentService2.checkout()

	MockPm := MockPaymentMethod{}
	paymentService3 := NewPaymentService(MockPm)
	paymentService3.checkout()
}
