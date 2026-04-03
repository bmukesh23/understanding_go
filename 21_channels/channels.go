package main

import "fmt"

// channels -> these are a way to communicate between goroutines.

// unbuffered channels: send and receive must happen at the same time.
// func main() {
// 	// define a channel of type string
// 	ch := make(chan string)

// 	go func() {
// 		// send a message to the channel
// 		ch <- "Hello, World!"
// 	}()

// 	// receive a message from the channel
// 	msg := <-ch // this waits until data is available in the channel
// 	fmt.Println(msg)
// }

// unbuffered channels, send and receive must run in different goroutines.
// new goroutine
//      |
//      |---- send (ch <- "Hello World")
// main goroutine
//      |
//      |---- receive  (<-ch)

// buffered channels: send and receive can happen at different times.
func main() {
	ch := make(chan int, 2) // buffer size of 2

	// send messages to the channel
	ch <- 1
	ch <- 2

	// receive messages from the channel
	fmt.Println(<-ch) // prints 1
	fmt.Println(<-ch) // prints 2
}

// closing channels: a channel can be closed when no more values will be sent.
// example:
/*
	ch := make(chan int)

	go func() {
     	ch <- 1
     	ch <- 2
     	close(ch)
	}()

 	for val := range ch {
    	fmt.Println(val)
	}
*/
