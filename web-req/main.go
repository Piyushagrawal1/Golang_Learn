package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning web request in Go")

	// -------- Make e GET Request------------

	response, err := http.Get("https://jsonplaceholder.typicode.com/users")
	//print type of response
	// fmt.Printf("Type of response: %T\n", response)
	//--------handle err------------
	if err != nil {
		fmt.Println("error makig GET Request: ", err)
		return
	}

	defer response.Body.Close()

	// -------Read the response body----------
	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		fmt.Println("error reading response: ", err)
		return
	}
	fmt.Println(string(data)) // data array of byte me tha isiliye use string me convert kiya
}
