package main

import "fmt"


type paymenter interface{
	pay(amount float32)
	refund(amount float32, accountNumber string)

}

type payment struct {
	gateway paymenter
	
	 
}

// open close principal

func (p payment) makePayment(amount float32) {
	//  razorpayPaymentGw := razorpay{}
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw.pay(amount)
	// stripePaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	//  logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

type fakepayment struct{}

func ( f fakepayment) pay(amount float32){
	fmt.Println("making payment using fakepayment", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }

type paypal struct{}

func (p paypal) pay(amount float32){
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, accountNumber string){
	fmt.Println("refund payment using paypal", amount, accountNumber)
}

func main() {
	// stripePaymentGw := stripe{}
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw,
	}
	newPayment.makePayment(100)
}
