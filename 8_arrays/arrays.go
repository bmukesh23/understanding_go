package main

import "fmt"

// array is a numbered sequence of specified length and element type.
func main() {
	// zeroed values for int
	// int -> 0, string -> "", bool -> false
	var nums [4]int

	// array length
	fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums)

	// to declare values inside the array in a single line
	num := [3]int{2, 4, 6}
	fmt.Println(num)

	// 2d array
	num2d := [2][2]int{{2, 4}, {3, 6}}
	fmt.Println(num2d)

	// fixed size, that is predictable
	// memory optimization
	// constant time access
}
