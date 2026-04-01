package main

import (
	"fmt"
	"sync"
)

// to run code in asynchronously or concurrently then we will use goroutines
// func task(id int) {
// 	fmt.Println("doing task", id)
// }

// goroutines -> main program will not wait for the goroutine to finish and it will continue executing the next lines of code.
// func main() {
// 	for i := 0; i <= 10; i++ {
// 		// go task(i)

// 		// we can also use anonymous function as goroutine
// 		go func(i int) {
// 			fmt.Println("doing task", i)
// 		}(i)
// 	}
// 	time.Sleep(time.Second * 1)
// }

// wait group -> It is a counter that tracks running goroutines and holds the main program until all of them finish.
func task(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Automatically called when function ends
	fmt.Println("doing task", id)
}
func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)       // Register one goroutine
		go task(i, &wg) // Start the goroutine
	}

	wg.Wait() // Wait here until all goroutines call Done()
}
