package main

import (
	"fmt"
	"strconv"
)

func main() {

	//--------Numeric type Conversion--------
	var num int = 42
	// fmt.Println("Number is: ", num)
	// fmt.Printf("Type of num is %T\n", num)

	var num2 float64 = float64(num)
	fmt.Println("Number is: ", num2)
	// fmt.Printf("Type of num is %T\n", num2)


	num = 123
	str := strconv.Itoa(num)
	fmt.Println("----Number is: ", str)
	fmt.Printf("-----Type of num is %T\n", str)

	num_str := "12345"
	number_int, _ := strconv.Atoi(num_str)
	fmt.Println("----Number is: ", number_int)
	fmt.Printf("-----Type of num is %T\n", number_int)
}