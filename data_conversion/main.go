package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num int = 42
	fmt.Println("number is ", num)
	fmt.Printf("type is %T\n", num)

	// num = num + 1.2 // here its gives error

	var data float64 = float64(num)
	data = data + 1.2
	fmt.Println("data is ", data)
	fmt.Printf("type is %T\n", data)

	num = 123
	str := strconv.Itoa(num)

	fmt.Println("value of str is ", str)
	fmt.Printf("type of str is  %T\n", str)

	number_str := "1234"
	number_int, _ := strconv.Atoi(number_str)

	fmt.Println("value of number_int is ", number_int)
	fmt.Printf("type of number_int is  %T\n", number_int)

	// converting string to float

	float_string1 := "3.14"
	number_float, _ := strconv.ParseFloat(float_string1, 64)
	fmt.Println("number float is ", number_float)
	fmt.Printf("type of number_float is  %T\n", number_float)

}
