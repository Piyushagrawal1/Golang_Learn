package main

import "fmt"

func modifiedValueByRefrence(num *int) {
	// *num = *num * 2
	*num *= 2
}

func main() {
	fmt.Println("Learning Pointer in Go")

	// var num int
	// num = 10

	// var ptr *int
	// ptr = &num

	num := "piyush"
	ptr := &num

	fmt.Println("num has value: ", num)
	fmt.Println("pointer contains: ", ptr)
	fmt.Println("Data contains through pointer: ", *ptr)


	value := 10
	modifiedValueByRefrence(&value)

	fmt.Println("value contains: ", value)
}