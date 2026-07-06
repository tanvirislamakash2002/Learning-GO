package main

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
