package main

import (
	"fmt"
	"mylearning/myutils"
)

func main() {
	fmt.Println("jai baaba ki ")

	var name string = "aashu "
	fmt.Println((name))

	var version = 1.0
	fmt.Println(version)

	var monney int = 50000
	fmt.Println(monney)

	var dimension float64 = 345.12
	fmt.Println(dimension)

	var decided bool = false
	decided = true
	fmt.Println(decided)

	const pi = 67.12
	fmt.Println(pi)

	// special type of varibale declarelation

	person := 123
	fmt.Println(person)

	// impoeting and exporting by capital letter
	// var Public = "data is impprtant "; // here the letter is capital
	var private = "data is private "

	fmt.Println((myutils.Public))
	fmt.Println((private))

}
