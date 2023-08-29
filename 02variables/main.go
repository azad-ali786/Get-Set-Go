package main

import "fmt"

// Globally declared variables should start with capital because its can access from outside | public
const JWTToken string = "dfsdfsdfsd"

func main() {
	var userName string = "Azad"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	var isLoggedIn bool = false;
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var valInt int
	fmt.Println(valInt)
	fmt.Printf("Variable is of type: %T \n", valInt)

	var valString string
	fmt.Println(valString)
	fmt.Printf("Variable is of type: %T \n", valString)

	// implicit type
	var x = "abc"
	fmt.Println(x)

	// no var style
	numberOfUser := 3000
	fmt.Println(numberOfUser)
}