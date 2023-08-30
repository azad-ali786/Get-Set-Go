package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct")
    
	azad := User {"Azad", "azadali78962@gmail.com", true, 16}
	fmt.Println("Azad", azad) // Azad {Azad azadali78962@gmail.com true 16}
	fmt.Printf("Azad details are: %+v \n", azad)  // Azad details are: {Name:Azad Email:azadali78962@gmail.com Status:true Age:16} 
}

type User struct {
	Name 	string
	Email 	string
	Status	bool
	Age		int				

}