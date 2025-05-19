package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:id`
	Title     string `json:"title"`
	Completed bool   `json:completed`
}

func main() {

	fmt.Println("learning crud ")

	res, error := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if error != nil {
		fmt.Println("having error 1", error)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Errorn in getting response ", res.Status)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if error != nil {
	// 	fmt.Println("having error 2", err)
	// 	return
	// }

	// fmt.Println("Data ", string(data))

	var todo Todo
	err := json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding :", err)
		return
	}
	fmt.Println("Todo :", todo)

}
