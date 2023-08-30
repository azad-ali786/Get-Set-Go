package main

import "fmt"

func main() {
	fmt.Println("Welcome to Slices")
    
	var sliceList = []string{"Apple","Tomato", "Peach"}
	fmt.Printf("Type of sliceList is %T \n", sliceList)

	sliceList = append(sliceList, "Mango")
	fmt.Println(sliceList)

	sliceList = append(sliceList[1:])
	fmt.Println(sliceList)

	// how to remove a value from slices based on index index
	var courses = []string{"reactjs","python", "swift", "ruby"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}