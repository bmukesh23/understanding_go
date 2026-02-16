package main

import "fmt"

// for -> only construct in Go for looping
func main() {
	// while loop
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	// classic for-loop
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	// range
	for i := range 10 {
		fmt.Println(i)
	}

	// continue and break
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}

		if i == 1 {
			break
		}

		fmt.Println(i)
	}
}
