package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the Files")
    content :="This is the content of the file"
	file, err := os.Create("./myFile.text")

	if err != nil {
		panic(err)
	} 

	length, err := io.WriteString(file,content)

	if err != nil {
		panic(err)
	} 

	fmt.Println("Length is :", length)

	readFile("./myFile.text")
	
	defer file.Close()
}


func readFile(fileName string) {
    dataByte,err := ioutil.ReadFile(fileName)
 
	if err != nil {
		panic(err)
	} 

	fmt.Println("Text data inside the files is \n", string(dataByte))


}