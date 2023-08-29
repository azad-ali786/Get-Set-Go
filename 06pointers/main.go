package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	num :=23

	var ptr = &num
	fmt.Println("Address of ptr", ptr)
	fmt.Println("Value of ptr", *ptr)

	*ptr = *ptr * 2
	fmt.Println(num)

}