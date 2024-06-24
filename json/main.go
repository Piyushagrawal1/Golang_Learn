package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("Learning JSON in GO")

	person := Person{
		Name:    "Piyush",
		Age:     21,
		IsAdult: true,
	}
	fmt.Println("Person Data is: ", person)

	//convert person in to JSON Encoding(Marshalling)

	personJSON, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshalling: ", err)
		return
	}

	fmt.Println("Person JSON data is: ", string(personJSON))


	//Decoding (UnMarshalling)

	var newPerson Person
	err = json.Unmarshal(personJSON, &newPerson)
	if err != nil {
		fmt.Println("Error UnMarshalling: ", err)
		return
	}
	fmt.Println("New Person Data is after unmarshalling: ", newPerson)

}
