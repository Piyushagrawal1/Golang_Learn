package main

import (
	"fmt"
)

func main() {

	//Switch with basics
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Not a valid day")
	}

	//Swich with multiple values
	month := "Fabuary"

	switch month {
	case "January", "Fabuary", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	case "July", "August", "September":
		fmt.Println("Summer")
	case "October", "November", "December":
		fmt.Println("Fall")
	default:
		fmt.Println("Not a valid month")
	}

	//Switch with Expressions
	temperature := 25

	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Very cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cold")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}
}
