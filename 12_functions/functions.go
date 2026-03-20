package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// function can return multiple values
func getData() (string, int, bool) {
	return "go", 33, true
}

// we can also return functions from functions
func getMultiplier() func(int) int {
	return func(x int) int {
		return x * 2
	}
}

func main() {
	// result := add(4, 6)
	// fmt.Println(result)

	// shorter way to write the above code
	fmt.Println(add(4, 6))

	// lang1, lang2, lang3 := getData()
	// fmt.Println(lang, num, flag)

	// shorter way to write the above code
	fmt.Println(getData())

	// using the function returned by getMultiplier
	multiplier := getMultiplier()
	fmt.Println(multiplier(5)) // Output: 10
}
