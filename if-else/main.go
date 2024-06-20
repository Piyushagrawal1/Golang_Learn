package main

import "fmt"

func main() {
	x := 15
	if x > 5 {
		fmt.Println("x is greater then 5")
	} else {
		fmt.Println("x is less then 5")
	}

	z := 20
	if z > 15 {
		fmt.Println("z is greater then 15")
	} else if z == 20 {
		fmt.Println("z is equals to 20")
	} else {
		fmt.Println("z is less then 5")
	}

	
}
