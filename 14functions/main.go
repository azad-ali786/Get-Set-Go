package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions")
    greeter()

	result := added(3,5)
	fmt.Println("Result of 3 plus 5 is", result)

	proResult, message := proAdded(3,2,5,4,7)
	fmt.Println(proResult)
	fmt.Println(message)

}

func proAdded( values ...int) (int, string) {
   total := 0

   for _,val := range values {
	total += val
   }

   return total, "Hi Pro result function"
}

func added(val1 int, val2 int) int {
	return val1 + val2
}
func greeter() {
	fmt.Println("Hello from Golang")
}