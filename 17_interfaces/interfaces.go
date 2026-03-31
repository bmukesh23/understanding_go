package main

import "fmt"

// using interface
// naming convention: ter at the end of the name, for example: pay -> paymenter, read -> reader, write -> writer
type paymenter interface {
	pay(amount float32)
}
type payment struct {
	gateaway paymenter
}

// we have to follow -> open close principle: open for extension but closed for modification
func (p payment) makePayment(amount float32) {
	// due to some updates we did modification in this makePayment method, which is not a good practice
	// razorPaymentGwt := razorPayment{}
	// razorPaymentGwt.pay(amount)

	// stripePaymentGwt := stripePayment{}
	// stripePaymentGwt.pay(amount)

	p.gateaway.pay(amount)
}

// type razorPayment struct{}
type stripePayment struct{}

// func (r razorPayment) pay(amount float32) {
// 	fmt.Println("making payment using razorpay", amount)
// }

func (s stripePayment) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

func main() {
	stripe := stripePayment{}
	np := payment{
		gateaway: stripe,
	}
	np.makePayment(100)
}
