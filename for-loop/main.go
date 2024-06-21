package main

import "fmt"

func main() {
	// BAsic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Number is: ",i)
	}

	//Infinite loop with break stmt

	counter := 0

	for {
		//infinite loop
		fmt.Println("Infinite Loop: ",counter)
		counter++

		//infinite loop stop krne k liye use Break Stmt
		if counter == 5 {
			break
		}
	}

	//range keyword

	numbers := []int {11,45,84,65,76,58}
	
	for index, value := range numbers {
		fmt.Println("Index is: ",index,"Value is: ",value)
	}

	data := "Piyush Agrawal"
	
	for index, char := range data {
		fmt.Println("Index is: ",index,"Value is: ",string(char))
	}
}