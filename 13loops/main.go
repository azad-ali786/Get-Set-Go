package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")
	
	days := []string{"Sun", "Mon","Tues", "Wed", "Thu", "Fri", "Sat"}

	for _, day := range days {
		fmt.Println("Days",day)
	}

	for d :=0; d< len(days); d++ {
		fmt.Println(days[d])
	}

	for index,day := range days {
		fmt.Printf("index is %v and values is %v \n",index,day)	
	}
    
	rogueValue :=1

	for rogueValue < 10 {
		fmt.Println("Value is", rogueValue)
		rogueValue++
	}
}