package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays")

	var fields [5]string
	fields[2] = "fff"

	fmt.Println("Fields", fields)
   

	var list =[2]string{"dfd", "dfdf"}
	fmt.Println("List", list)
}