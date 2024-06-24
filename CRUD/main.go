package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"net/http"
)

// create struct
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetResponse() {
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

func performPostResponse() {
	todo := Todo{
		UserID:    21,
		Title:     "Learn Go",
		Completed: true,
	}
	//convert the todo struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error Marshalling: ", err)
		return
	}
	// fmt.Println("jsonData: ", string(jsonData))

	//convert string data to reader
	stringData := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(stringData)

	resp, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error Posting: ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status response: ", resp.Status) //201 Created

	data, _ := io.ReadAll(resp.Body)
	fmt.Println("Data: ", string(data))
}

func performUpdateResponse() {
	todo := Todo{
		UserID:    21852,
		Title:     "Learn Go with Piyush Agrawal",
		Completed: false,
	}

	data, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error Marshalling: ", err)
		return
	}

	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	//convert string data to reader
	stringData := string(data)

	//convert string data to reader
	jsonReader := strings.NewReader(stringData)

	req, err := http.NewRequest("PUT", myUrl, jsonReader)
	if err != nil {
		fmt.Println("Error Creating: ", err)
		return
	}
	req.Header.Set("content-type", "application/json")

	//send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Sending: ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status response: ", resp.Status)

	datas, _ := io.ReadAll(resp.Body)
	fmt.Println("Data: ", string(datas))

}

func main() {
	fmt.Println("Learning CRUD in GO")

	performGetResponse()

	performPostResponse()

	performUpdateResponse()
}
