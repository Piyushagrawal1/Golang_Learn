package main

import "fmt"

func main() {

	// name <--> grade
	studentGrade := make(map[string]int) //Syntex

	studentGrade["Piyush"] = 100
	studentGrade["Mohit"] = 80
	studentGrade["Yash"] = 60
	studentGrade["Nikhil"] = 40

	//print marks of Yash
	fmt.Println("Marks of Yash: ", studentGrade["Yash"])
	//change marks of Yash
	studentGrade["Yash"] = 70
	fmt.Println("marks changed")
	//print new marks
	fmt.Println("New Marks of Yash: ", studentGrade["Yash"])

	//delete Yash
	// delete(studentGrade, "Yash")
	fmt.Println("Marks of Yash: ", studentGrade["Yash"])

	//check if a key exist or not
	grades, exist := studentGrade["Nishant"] // grades ==> value, exist ==> boolian

	fmt.Println("Grades of Nishant: ", grades)
	fmt.Println("Existance of Nishant: ", exist)

	Grades, Exist := studentGrade["Mohit"] // grades ==> value, exist ==> boolian

	fmt.Println("Grades of Mohit: ", Grades)
	fmt.Println("Existance of Mohit: ", Exist)

	//use range to iterate all
	for index, value := range studentGrade {
		// fmt.Println("Key is: ", index, "Marks is: ", value)
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	//another way of initialization

	person := map[string]int {
		"Piyush": 100,
		"Mohit":  80,
		"Yash":   60,
		"Nikhil": 40,
	}
	for index, value := range person {
		fmt.Printf("-----------Key is %s and marks is %d\n", index, value)
	}
}
