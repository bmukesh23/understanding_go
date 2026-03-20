package main

import "fmt"

// variadic functions are functions that can take a variable number of arguments.
// This is achieved by using an ellipsis(...) before the type of the last parameter in the function definition.

// rest operator
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	// result := sum(2, 3, 5)

	nums := []int{2, 3, 5}
	// spread operator
	result := sum(nums...)
	fmt.Println(result)
}
