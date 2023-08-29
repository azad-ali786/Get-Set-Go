package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Times in golang")

	presentTime := time.Now();
	fmt.Println("The present time is", presentTime)
	fmt.Println("The formatted present time is", presentTime.Format("01-02-2006 Monday"))

	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println("Created date", createdDate)
	fmt.Println("The formatted present time is", createdDate.Format("01-02-2006 Monday"))
}