package main

import "fmt"

func main() {
	// simple switch
	i := 3

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	// default is optional
	default:
		fmt.Println("other")
	}

	// multiple condition switch

}
