package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {

	fmt.Println("Starting of execution")

	data := add(2, 3)
	defer fmt.Println("data is: ", data)

	// fmt.Println(add(1, 2))
	defer fmt.Println("middle of execution") //defer vala sabse end me print hoga
	fmt.Println("ending of execution")

}

//Note: agar do jagah defer laga huaa hai to vo LIFO condition ko follow krta hai
//first me lga huaa defer vala sabse end me print hoga
