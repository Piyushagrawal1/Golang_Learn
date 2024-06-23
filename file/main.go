package main

import (
	"fmt"
	"os"
)

func main() {

	/*//file creation
	file, err := os.Create("piyush.txt")
	if err != nil {
		fmt.Println("Error wile createing file: ", err)
		return
	}

	defer file.Close()


	//-------1st way:- writing to file------

	content := "ha bhai kya hal hai"
	_, error := io.WriteString(file, content)

	if error != nil {
		fmt.Println("Error while writting file: ", error)
		return
	}

	fmt.Println("File created successfully")*/

	//------------another way:- writing to file------------
	// fmt.Println("File created successfully")
	// file.WriteString("Hello, World!")

	/*
		//----------Reading file Buffer---------------

		file, err := os.Open("piyush.txt")
		if err != nil {
			fmt.Println("Error while opening file: ", err)
			return
		}

		defer file.Close()

		//---------------------creating a buffer to read file content----------------

		text := make([]byte, 100)
		len, err := file.Read(text)
		if err != nil {
			fmt.Println("Error while reading file: ", err)
			return
		}
		fmt.Println("File content: ", string(text[:len]))*/

	//------------------ Read file using ioutil function but ioutil is depreciated so we use os

	content, err := os.ReadFile("piyush.txt")

	if err != nil {
		fmt.Println("Error while reading file: ", err)
		return
	}

	fmt.Println("File content: ", string(content))
}
