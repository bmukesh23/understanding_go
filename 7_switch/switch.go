package main

import (
	"fmt"
	"time"
)

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
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's a weekend")
	default:
		fmt.Println("it's workday")
	}

	// type switch
	whoAmI := func(i any) {
		switch i.(type) {
		case int:
			fmt.Println("it's an int")
		case string:
			fmt.Println("it's a string")
		default:
			fmt.Printf("other")
		}
	}

	whoAmI(50)
	whoAmI("golang")
	whoAmI(22.5)
}
