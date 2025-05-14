package main

import "fmt"

func main() {
	fmt.Println("we are studying if else ")

	x := 10
	if x > 5 {
		fmt.Println("value is 5")
	} else {
		fmt.Println("value is not 5")
	}

	// if else

	z := 10
	if z > 10 {
		fmt.Println("thr value is greater than 10")
	} else if z > 5 {
		fmt.Println("the value is greater than 5")

	} else {
		fmt.Println("value is less than 5")
	}

}
