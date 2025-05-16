package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/")

	if err != nil {
		fmt.Println("Error in getting the api data ", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("tType of response is %T\n", res)
	// fmt.Println("data us ", res)

	// reading fro response

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error in reading the data")
		return
	}

	fmt.Println("respns eis ", string(data))

}
