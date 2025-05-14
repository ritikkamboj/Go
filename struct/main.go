package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	fmt.Println("learning the struct in go ")

	var Prince Person
	fmt.Println(" Prince is ", Prince)
	Prince.firstName = "prince"
	Prince.lastName = "kamboj"
	Prince.age = 12

	fmt.Println(" Prince is ", Prince)

	// defining in other way

	Person1 := Person{
		firstName: "Ritik",
		lastName:  "Kamboj",
		age:       26,
	}

	fmt.Println("the value of Person1 is ", Person1)

	// 3rd way

	var person2 = new(Person)
	person2.firstName = "Navdeep"
	person2.lastName = "singh"
	person2.age = 25

	fmt.Println("the value of navdeep is ", person2)

	type Contact struct {
		email string
		phone int
	}
	type Address struct {
		house int
		area  string
		state string
	}

	type Employe struct {
		person_details Person
		person_contact Contact
		person_address Address
	}

	var employe1 Employe
	employe1.person_details = Person{
		firstName: "priya",
		lastName:  "Kamboj",
		age:       29,
	}
	employe1.person_contact.email = "abc@gmail.com"
	employe1.person_contact.phone = 12314567890

	employe1.person_address = Address{
		house: 12,
		area:  "ranchi",
		state: "Haryana",
	}
	fmt.Println("the value of nested structure is ", employe1)
}
