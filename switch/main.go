package main

import (
	"fmt"
)

func main() {
	fmt.Println("studying about the switch case ")

	name := 2

	switch name {
	case 1:
		fmt.Println("the value is one")
	default:
		fmt.Println("the value is not one ")
	}

	// using to make program on weather

	month := "january"

	switch month {
	case "january", "february", "december":
		fmt.Println("its winter time")
	case "may", "june", "july":
		fmt.Println("its peak summer time")
	default:
		fmt.Println("no weather ")
	}

	// using conditionall in switch

	temperatur := 25

	switch {
	case temperatur < 0:
		fmt.Println("freezing")
	case temperatur >= 0 && temperatur < 10:
		fmt.Println("cold")
	case temperatur >= 10 && temperatur < 25:
		fmt.Println("cool")
	case temperatur >= 25 && temperatur < 30:
		fmt.Println("hot")
	default:
		fmt.Println("so hot")
	}

}
