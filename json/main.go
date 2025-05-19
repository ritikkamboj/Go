package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {

	person := Person{
		Name:    "ritik",
		Age:     23,
		IsAdult: true,
	}

	fmt.Println("person is ", person)

	// converting in json

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Entering marshalling ", err)
		return
	}
	fmt.Println("person data is ", string(jsonData))

	// decoding and unmarshalling

	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error Unmarshalling", err)
		return
	}

	fmt.Println("Person data is", personData)

}
