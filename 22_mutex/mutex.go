package main

import (
	"fmt"
	"sync"
)

// mutex -> is also known as mutual exclusion lock.
// it is used to prevent race conditions when multiple goroutines access shared data concurrently.
// it is a lock that ensures that only one goroutine accesses shared data at a time.
var counter int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	mu.Lock() // lock
	counter++
	mu.Unlock() // unlock
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println(counter)
}
