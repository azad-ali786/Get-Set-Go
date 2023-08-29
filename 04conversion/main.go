package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to Conversions"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the code:")
    
	// comma ok || error ok
	input, _ := reader.ReadString('\n')

	fmt.Println("The rating is",input) 

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added one to rating",numRating + 1)
	}

}