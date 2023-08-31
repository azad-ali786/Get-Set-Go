package main

import "fmt"

func main() {
	fmt.Println("Welcome to Defer")
	defer fmt.Println("HI I am deferred 1")
	defer fmt.Println("HI I am deferred 2")
	fmt.Println("Oh! I am executed first")

}

// Output :
// Welcome to Defer
// Oh! I am executed first
// HI I am deferred 2
// HI I am deferred 1


// Defer --> Before Return and LIFO