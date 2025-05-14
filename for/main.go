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

	// range keyword
	// use with slice and string

	numbers := []int{12, 45, 67, 15, 64}

	for index, value := range numbers {
		fmt.Printf("index of the number is %d and value is %d\n", index, value)
	}

	// use this with string case

	data := "jai maata di"
	for index, value := range data {
		fmt.Printf("index is %d and character is %c\n", index, value)
	}
}
