package main

import "fmt"

func main() {
	var name string = "mukesh"
	fmt.Println(name)

	var age int = 30
	fmt.Println(age)

	// infer
	var isMarried = true
	fmt.Println(isMarried)

	// shorthand syntax
	name2 := "mukesh"
	fmt.Println(name2)

	// variable declaration and value assignment can be done seperately in "var" not in ":="
	var price float32
	price = 50.5
	fmt.Println(price)
}
