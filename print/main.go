package main

import "fmt"

func main() {
	age := 45
	name := "Piyush"
	height := 6.264545124

	fmt.Println("age: ", age, "Height: ", height, "Name: ", name)
	fmt.Println("hello world")

	fmt.Printf("age is %d\n", age)
	fmt.Printf("Height is %.3f\n", height)
	//find the data type 
	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Type of height is %T\n", height)

	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)

}