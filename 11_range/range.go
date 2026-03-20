package main

import "fmt"

// iterating over data structures
func main() {
	// nums := []int{2, 4, 6}

	// normal for loop
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// using range
	// for _, num := range nums {
	// 	fmt.Println(num)
	// }

	// sum of numbers using range
	// sum := 0
	// for _, num := range nums {
	// 	sum = sum + num
	// }
	// fmt.Println(sum)

	// i is index in this range loop
	// for i, num := range nums {
	// 	fmt.Println(i, num)
	// }

	// range can also be used with maps
	// m := map[string]string{"fname": "John", "lname": "Doe"}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// range can also be used with strings
	for i, c := range "Hello" {
		// it prints the c -> in unicode code point rune
		// in this i is -> starting byte of rune
		// fmt.Println(i, c)

		// to print character instead of unicode code point
		fmt.Println(i, string(c))
	}
}
