package main

import "fmt"

type person struct {
	FirstName string
	Lastname  string
	Age       int
}

type Contact struct {
	Phone string
	Email string
}

type Address struct {
	House int
	Area string
	Conutry string
}

type Employee struct {
	Person_Details person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	//------1st way to use struct---------
	var piyush person
	// fmt.Println("Piyush person: ", piyush) // " ", " ", 0 // __0

	piyush.FirstName = "Piyush"
	piyush.Lastname = "Agrawal"
	piyush.Age = 21
	// fmt.Println("Piyush person: ", piyush)

	// ----------2nd way to use struct---------
	person1 := person{
		FirstName: "Akash",
		Lastname:  "Agrawal",
		Age:       28,
	}
	fmt.Println("Piyush person: ", person1)

	// ----------3rd way to use struct with new keyword---------
	var person2 = new(person)

	person2.FirstName = "Muskan"
	person2.Lastname = "Agrawal"
	person2.Age = 18

	// fmt.Println("Muskan person: ", person2)
	// fmt.Println("Age of Muskan: ", person2.Age)

	var employee1 Employee

	// employee1.Person_Details = *person2 

	employee1.Person_Details = person {   //ye vala tarika upar vale tarike ko overlap kr deta hai
		FirstName: "Piyush",
		Lastname:  "Agrawal",
		Age:       22,
	}

	employee1.Person_Contact.Email = "piyush@gmail.com"
	employee1.Person_Contact.Phone = "1234567890"

	employee1.Person_Address = Address{
		House: 123,
		Area:  "pune",
		Conutry:  "pune",
	}

	fmt.Println("Employee 1 : ", employee1)

	fmt.Println("Name: ", employee1.Person_Contact.Email)

}
