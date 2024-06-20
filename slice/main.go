package main

import "fmt"

func main() {
	fmt.Println("Slice in Go")

	// number := []int {1,2,3,4,5}
	// //append
	// number = append(number, 6,7,8,9,10,11,12,13,14,15)
	// //print the number
	// fmt.Println("Number is: ", number)
	// //find the type
	// fmt.Printf("number has datatype: %T\n", number)
	// //find the length
	// fmt.Println("length is: ", len(number))

	// number := []int{1, 2, 3}
	// fmt.Println("slice:", number)
	// fmt.Println("length:", len(number))
	// fmt.Println("capacity:", cap(number))

	//-------make-----
	number := make([]int, 3, 5) // by default stores in 0
	number = append(number, 4, 98, 45)
	fmt.Println("slice:", number)
	fmt.Println("length:", len(number))
	fmt.Println("capacity:", cap(number))
}
