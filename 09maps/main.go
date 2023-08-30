package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["ts"] = "typescript"
	languages["rb"] = "ruby"

	fmt.Println("List of all languages", languages)

	fmt.Println("JS: ",languages["js"])

	delete(languages, "rb")

	fmt.Println("List of all languages", languages)

   // loops 
    for key,value := range languages {
		fmt.Printf("For key %v, value %v \n",key,value)
	}
}