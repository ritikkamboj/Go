package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// crating file

	/*
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("there is some error in file creating ")
			return
		}

		content := "hello World "

		_, err1 := io.WriteString(file, content+"\n")

		if err1 != nil {
			fmt.Println("error while entering some data in file ")

		}
	*/

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("there is some error in file creating ")
		return
	}

	defer file.Close()

	buffer := make([]byte, 1024)

	// reading the file

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error in reading the file ", err)
			return

		}
		// process the read from buffer
		fmt.Println(string(buffer[:n]))
	}

	// another method when more data then not recommended to use
	//ioutil

	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("error in reading the file ", err)
		return

	}

	fmt.Println(string(content))

}
