package main

import (
	"fmt"
	"net/url"
)

const targetUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=lkjdsfl334kj"

func main() {
	fmt.Println("Welcome to the URL handling")
  
	result, _  := url.Parse(targetUrl)
    
	fmt.Println("Result scheme",result.Scheme)
	fmt.Println("Result HOst",result.Host)
	fmt.Println("Result Path",result.Path)
	fmt.Println("Result Query",result.RawQuery)
	fmt.Println("Result Port",result.Port())

	qparams := result.Query();
	fmt.Println(qparams)

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "azad",
		Path: "sus",
		RawQuery: "user=me",
	}
	// &url.URL is going to create a pointer, To modify the URL later and use the updated value we need to use pointers. If we directly use the URL then if the URL gets changed later, it will be in a new copy of the URL instead of the original URL being modified.

	fmt.Println("Newly created URL", partsOfUrl.String())
}

