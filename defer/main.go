package main

import "fmt"

func main() {

	fmt.Println("starting of the project ")
	defer fmt.Println("middle of the project ")
	fmt.Println("end of the project ")

	// defer follow stack when we use more than one defer
}
