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
	EncodeJson()

}

func EncodeJson() {
	courses := []course{
		{"Reactjs", 299, "something.in", "abc123", []string{"webdev","js"}},
		{"Vuejs", 199, "something.in", "xyz123", nil},
		{"MERN", 199, "something.in", "fsdf3", []string{"webdev","ts","fullstack"}},
	}

	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err!=nil {
		panic(err)
	}

	fmt.Printf("JSON", string(finalJson))

}