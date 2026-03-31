package main

import "fmt"

// generics where introduced in 1.18 version

// func printSliceInt(items []int) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
// func printSliceString(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// to perform multiple type operations without writing the same code for different types, we use generics
// if it is difficult to write all types in the generic function then use comparable -> func printSlice[T comparable ] (items []T){}
// we can also perfor multiple generics in the same function -> func printSlice[T int | string | bool, U float64 ] (items []T, value U){}
func printSlice[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// we can also use generics with struct
type stack[T int | string | bool] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	strs := []string{"go", "c++", "java"}
	vars := []bool{true, false}

	printSlice(nums)
	printSlice(strs)
	printSlice(vars)

	myStack := stack[int]{
		elements: []int{1, 2, 3, 4, 5},
	}

	fmt.Println(myStack)
}
