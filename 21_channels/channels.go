package main

import "fmt"

// channels are a way to communicate between goroutines. They are a powerful tool for concurrent programming in Go. Channels can be used to send and receive values between goroutines, and they can also be used to synchronize the execution of goroutines.
func main() {
	// define a channel of type string
	ch := make(chan string)

	go func() {
		// send a message to the channel
		ch <- "Hello, World!"
	}()

	// receive a message from the channel
	msg := <-ch // this waits until data is available in the channel
	fmt.Println(msg)
}
