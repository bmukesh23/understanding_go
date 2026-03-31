package main

import "fmt"

// go doesn't have built-in support for enums, but we can use const to achieve similar functionality

// enumerated types
// int example
type orderStatusInInt int

const (
	// iota is a untyped integer and special identifier that is reset to 0.
	Receive orderStatusInInt = iota //0
	Process                         //1
	Ship                            //2
	Deliver                         //3
)

// string example
type orderStatusInString string

const (
	Received   orderStatusInString = "received"
	Processing orderStatusInString = "processing"
	Shipped    orderStatusInString = "shipped"
	Delivered  orderStatusInString = "delivered"
)

func changeOrderStatus(status orderStatusInString) {
	fmt.Println("Order status changed to", status)
}

func main() {
	changeOrderStatus(Received)
}
