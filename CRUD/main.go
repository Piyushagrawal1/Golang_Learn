package main

import (
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"net/http"
)

// create struct
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD in GO")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error Getting: ", err)
		return
	}

	defer resp.Body.Close()

	// fmt.Println("Status Code: ", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error Response: ", resp.StatusCode)
		return
	}

	//  ---------is method se directly json data print kr diya pan ne-------------//

	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error Reading: ", err)
	// 	return
	// }

	// // fmt.Println("Data: ", string(data))

	// var todo Todo
	// err = json.Unmarshal(data, &todo)
	// if err != nil {
	// 	fmt.Println("Error Decoding: ", err)
	// 	return
	// }

	// fmt.Println("Todo: ", todo)

	//-----------is se hame unmarshal krne ki need nahi hai directly uske jaisa o/p de dega-----

	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error Decoding: ", err)
		return
	}

	fmt.Println("Todo: ", todo)
	

}
