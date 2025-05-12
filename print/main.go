package main

import "fmt"

func main() {
	age := 25
	name := "ritik"
	height := 5.87651616516

	fmt.Println("age: ", age, "height: ", height, "name: ", name)
	fmt.Printf("age is %d", age)
	fmt.Printf("height is %.4f", height)
	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("name is  %s\n", name)

}
