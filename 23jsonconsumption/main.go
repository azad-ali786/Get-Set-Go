package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"courseName"`
	Price int `json:"price"`
	Platform string `json:"website"`
	Password string `json:"-"` // dash signifies that whoever is consuming my api will not see password
	Tags []string `json:"tags,omitempty"` // omitempty says if the value is null, the field wont be visible Note: don't keep empty space before omitempty
}

func main() {
	fmt.Println("Welcome to Json Handling")
	DecodeJson()

}

func DecodeJson() {
	jsonData := []byte(`{
                "courseName": "MERN",
                "price": 199,
                "website": "something.in",
                "tags": ["webdev","ts","fullstack"]
        } `)
	
		var courses course

		checkValid := json.Valid(jsonData)

		if checkValid {
			fmt.Println("It is valid")
			json.Unmarshal(jsonData, &courses)
			fmt.Printf("%#v\n", courses)
		} else {
			 fmt.Println("It is not valid")
		}

		// extract based on key value 
		var coursesData map[string] interface{}
		json.Unmarshal(jsonData, &coursesData)
		fmt.Printf("%#v\n", coursesData)

		for k,v := range coursesData {
			fmt.Printf("key is %v and value is %v \n", k, v)
		}
}