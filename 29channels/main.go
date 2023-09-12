package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to Channels")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)   // deadlock 

	wg.Add(2)
   // Recieve only because of <-ch , cant close channel here
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChOpen := <-myCh

		if isChOpen {
        fmt.Println(val)
		}

		wg.Done()
	}(myCh, wg)


	// Send only because of <-ch
	go func(ch chan<- int, wg *sync.WaitGroup) {
		close(ch)
		wg.Done()
	}(myCh, wg)

    wg.Wait()
}

// Note: Listening on a close channel is allowed but writing is not
// if channel is close, it returns 0  