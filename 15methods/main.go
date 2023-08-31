package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct")
    
	azad := User {"Azad", "azadali78962@gmail.com", true, 16}
	fmt.Println("Azad", azad) // Azad {Azad azadali78962@gmail.com true 16}
	fmt.Printf("Azad details are: %+v \n", azad)  // Azad details are: {Name:Azad Email:azadali78962@gmail.com Status:true Age:16} 
    fmt.Println(azad.GetStatus())
	azad.NewMail()
	fmt.Println("New email of the User is", azad.Email)  // using NewMail function does not change the respect mail, because the object is not passed by reference
}


type User struct {
	Name 	string
	Email 	string
	Status	bool
	Age		int				
}

func (u User) GetStatus() bool {
   return u.Status;
}

func (u User) NewMail() {
	u.Email = "azadxyz@gmail.com"
}