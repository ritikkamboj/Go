package main

import (
	"fmt"
)

func main() {
	fmt.Println("seeing the for loop")

	for i := 0; i < 10; i++ {
		fmt.Println("numbers are ", i)

	}

	// infinite loop in go

	counter := 0

	for {

		fmt.Println("Infinite Loop")
		counter++
		if counter == 3 {
			break
		}
	}

}
