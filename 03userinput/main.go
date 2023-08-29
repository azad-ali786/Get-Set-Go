package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user inputs"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the code:")
    
	// comma ok || error ok
	input, _ := reader.ReadString('\n')

	fmt.Println("The rating is",input) 

}