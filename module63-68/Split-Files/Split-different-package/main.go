package main

import "learnGO/module63-68/Split-Files/Split-different-package/payment"

func main() {
	// bkash := Bkash{apiKey: "4343hjkjj"}
	// paymentService1 := NewPaymentService(&bkash)
	// paymentService1.checkout()

	// nagad := NewNagad("ksddff83in")
	// paymentService2 := NewPaymentService(nagad)
	// paymentService2.checkout()

	MockPm := payment.MockPaymentMethod{}
	paymentService3 := payment.NewPaymentService(MockPm)
	paymentService3.Checkout()
}
