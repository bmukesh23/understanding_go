package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	// creating map

	m := make(map[string]string)

	// adding key value pair
	m["name"] = "golang"
	m["age"] = "fifteen"

	// get an element
	fmt.Println(m["name"])

	// get length of map
	fmt.Println(len(m))

	// delete an element
	delete(m, "age")
	fmt.Println(m)

	// clear map -- these are the built-in methods
	clear(m)

	// another way to create map
	mp := map[string]int{"age": 15, "year": 2024}
	fmt.Println(mp)

	// we use this for checking if a key exists in the map
	// ok -> returns boolean and _ -> returns value
	_, ok := mp["age"]

	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	// maps.Equal() -> checks if two maps are equal
	m1 := map[string]int{"price": 40, "phones": 30}
	m2 := map[string]int{"price": 40, "phones": 30}

	fmt.Println(maps.Equal(m1, m2))
}
