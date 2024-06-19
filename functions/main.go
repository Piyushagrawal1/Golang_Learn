package main

import "fmt"

func simpleFunction() {
	fmt.Println("Simple Function")
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b, c int) int {
	return a * b * c
}

func subtract(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("We are learning function in Golang")
	simpleFunction()

	ans := add(3, 4)
	fmt.Println("Addition is: ", ans)

	data := multiply(3, 4, 5)
	fmt.Println("Multiplication is: ", data)

	data = subtract(4, 3)
	fmt.Println("Subtract is: ", data)
}
