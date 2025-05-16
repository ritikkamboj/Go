package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("urls ")

	myurl := "https://jsonplaceholder.typicode.com/todos/1"

	parseUrl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("there is some error ")
		return
	}
	fmt.Printf("type of the url is %T\n", parseUrl)
	fmt.Println("schema is ", parseUrl.Scheme)
	fmt.Println("apth is ", parseUrl.Path)
	fmt.Println("Raw query ", parseUrl.RawQuery)
	fmt.Println("Host name", parseUrl.Host)

	// Modefying the components

	parseUrl.Path = "/jaishreeram"
	parseUrl.RawQuery = "username= prince "

	//  creating a new url

	newUrl := parseUrl.String()
	fmt.Println("new URl us ", newUrl)

}
