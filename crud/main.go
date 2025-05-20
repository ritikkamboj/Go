package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:id`
	Title     string `json:"title"`
	Completed bool   `json:completed`
}

func performGetRequest() {
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

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Ritik kamboj",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in marshalling", err)
		return
	}

	//convert ot string
	jsonString := string(jsonData)

	// converting the string data to reader
	jsonReader := strings.NewReader(jsonString)
	res, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending in data", err)
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response is ", string(data))

}

func performUpdateRequest() {
	todo := Todo{
		UserID:    232121212,
		Title:     "Ritik kamboj know golang ver well ",
		Completed: false,
	}
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in marshalling", err)
		return
	}

	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	//convert ot string
	jsonString := string(jsonData)

	// converting the string data to reader
	jsonReader := strings.NewReader(jsonString)

	// for doing put request we don't have a direct method for that

	req, err := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err != nil {
		fmt.Println("error ub putting ", err)
		return
	}
	req.Header.Set("Content-type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sendind request ", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))

}

func main() {

	fmt.Println("learning crud ")
	// performGetRequest()
	// performPostRequest()

	performUpdateRequest()
}
