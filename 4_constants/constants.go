package main

import "fmt"

const isEligible = true

func main() {
	// const name = "golang"
	// const age = 35

	// ":=" can only be used inside the function.

	fmt.Println(isEligible)

	// Multiple constants can be declared together using a single const keyword and parentheses.
	const (
		name = "golang"
		age  = 35
		num  = 50.1
	)
}
