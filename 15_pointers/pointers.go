package main

import "fmt"

// go passes arguments by value, so when we pass num to changeNum, it creates a copy of num. Therefore, when we change num inside changeNum, it does not affect the original num in main. The output will be:
// In changeNum 5
// After changeNum in main 1

// by value
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

// by reference
// num *int means num is a pointer to an int.
func changeNum(num *int) {
	// *num means value at pointer
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1
	changeNum(&num)
	fmt.Println("After changeNum in main", num)

	// To see memory addresses, we can use the & operator to get the address of a variable.
	// fmt.Println("memory address", &num)
}
