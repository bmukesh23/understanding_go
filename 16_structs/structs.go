package main

import (
	"fmt"
	"time"
)

type order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time //nanosecond precision
}

// receiver type
// func (receiver Type) methodName(parameters) {}
func (o *order) changeStatus(status string) {
	// struct does dereference automatically, so you can use o.status instead of (*o).status
	o.status = status
}

func main() {
	// if you dont set any field, default value is zero value
	// int, float -> 0, string -> "", bool -> false
	myOrder := order{
		id:     1,
		amount: 50.00,
		status: "received",
	}

	myOrder.createdAt = time.Now()
	// fmt.Println("order struct id:", myOrder.id)

	myOrder2 := order{
		id:        2,
		amount:    130.00,
		status:    "pending",
		createdAt: time.Now(),
	}

	myOrder.status = "shipped"

	// struct won't interfere with each other, they are different instances of the same struct type
	// fmt.Println("order struct", myOrder)
	fmt.Println("order struct", myOrder2)

	myOrder.changeStatus("paid")
	fmt.Println("order struct", myOrder)

	// we can directly create a struct without assigning it to a variable
	language := struct {
		name string
		year int
	}{"Go", 2009}
	fmt.Println("language struct", language)
}

// there is an another concept known as struct embedding, which allows you to include one struct within another struct.

// type customer struct {
// 	name  string
// 	email string
// }

// type orderedCustomer struct {
// 	id        int
// 	amount float32
// 	customer // embedding the customer struct
// }

// func main(){
// myCustomer := customer{
// 		name: "mukesh"
// 		email: "mukesh123@gmail.com"
// }

// myOrder := order{
// 		id:     1,
// 		amount: 50.00,
// 		status: "received",
// 		customer: myCustomer
// }
// }
