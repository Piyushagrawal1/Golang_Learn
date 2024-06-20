package main

import "fmt"

func main() {
	fmt.Println("Array in GO")

	// array Declaration
	var name [5]string

	name[0] = "Piyush"
	name[2] = "Agrawal"

	fmt.Println("Name of person is: ", name)

	//array initialization
	var marks = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Marks of person is: ", marks)

	// find length of an array
	fmt.Println("Length of array is: ", len(marks))

	//find the value in index
	fmt.Println("value at int 2 is: ", name[2])



}
