package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL handling in Go")
	myUrl := "https://www.google.com/search?q=piyush+agrawal"
	// fmt.Printf("Type of URL is: %T\n",  myUrl)

	parsedURl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("can't parse url", myUrl)
		return
	}
	fmt.Printf("Type uf URL: %T\n", parsedURl)

	fmt.Println("Scheme: ",parsedURl.Scheme)
	fmt.Println("Host: ",parsedURl.Host)
	fmt.Println("Path: ",parsedURl.Path)
	fmt.Println("RawQuery: ",parsedURl.RawQuery)

	//modyfy url
	parsedURl.Path = "/api/v1/search"
	
	newPath := parsedURl.String()
	fmt.Println("New url: ",newPath)
}
