package main

import "fmt"

func main() {
	age := 2

	// if-else statement
	if age >= 18 {
		fmt.Println("person is an adult")
	} else {
		fmt.Println("person is not an adult")
	}

	// else-if statement
	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is a teenager")
	} else {
		fmt.Println("person is a kid")
	}

	// logical operators
	var role = "admin"
	var hasPermission = false

	// if role == "admin" || hasPermission {
	// 	fmt.Println("yes")
	// }

	if role == "admin" && hasPermission {
		fmt.Println("yes")
	}

	// we can declare a variable in the if statement
	if age := 15; age >= 18 {
		fmt.Println("person is an adult", age)
	} else {
		fmt.Println("person is not an adult", age)
	}

	// Go does not have ternary operator, you will have to use normal if-else
}
