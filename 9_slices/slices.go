package main

import (
	"fmt"
	"slices"
)

//	slices -> dynamic arrays
//
// most used construct in go
// plus point that it as a useful methods
func main() {
	// unintialized slice is nil (null in other languages)
	var nums []int
	fmt.Println(nums == nil) // true

	var num = make([]int, 2)
	fmt.Println(num == nil) // false

	// capacity -> maximum no. of elements can fit
	fmt.Println(cap(num)) // 2

	var num2 = make([]int, 0, 3)
	// num2 := []int{}
	num2 = append(num2, 1)
	num2 = append(num2, 2)
	// num2[0] = 1
	// num2[1] = 2

	fmt.Println(num2)
	fmt.Println(cap(num2))

	// copy
	var val = make([]int, 0, 2)
	val = append(val, 1)
	var val2 = make([]int, len(val))
	copy(val2, val)
	fmt.Println(val, val2)

	// slice operator
	var num3 = []int{1, 2, 3}
	fmt.Println(num3[0:2]) // [1 2]
	fmt.Println(num3[:2])  // [1 2]
	fmt.Println(num3[1:])  // [2 3]

	// slice -> there is another package called slices which has more methods for slices. it's an in-built package
	var num4 = []int{1, 2, 3}
	var num5 = []int{4, 5, 6}
	fmt.Println(slices.Equal(num4, num5)) //false

	// 2d slices
	var matrix = [][]int{{1, 3, 4}, {6, 4, 3}}
	fmt.Println(matrix)
}
