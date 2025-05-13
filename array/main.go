package main

import "fmt"

func main() {
	fmt.Println("We are learning array ")

	var name [5]string
	name[0] = "prie"
	name[2] = "aashu"

	fmt.Println("Names of the Person is", name)

	// array initialization
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Number is ", numbers)

	// to find the length of array

	fmt.Println("Length of Numbers array is :", len(numbers))

	// when we have to find the index value

	fmt.Println("value at index 2 is ", numbers[2])

	var price [5]int
	fmt.Println("price is ", price)

	var name1 [5]string
	name1[1] = "ritik"
	name1[2] = "aakash"

	fmt.Printf("name is : %q", name1)

}
