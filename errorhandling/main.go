package main

import "fmt"

// error handling in function
// func division(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("denominator can't be zero ")

// 	}

// 	return a / b, nil

// 	// return
// }

func division(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "denominator can't be zero "

	}

	return a / b, "nil"

	// return
}

func main() {
	fmt.Println("we are seeing error hadnling ")

	res, _ := division(10, 2)
	fmt.Println("The value of division is", res)
}
