package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello world")
	// time.Sleep(2000 * time.Millisecond) // stimulating some work
	fmt.Println("Hello world ended")

}

func sayHi() {
	fmt.Println("Hi Prince ")
	time.Sleep(1000 * time.Millisecond) // stimulating some work

}

func main() {
	fmt.Println("learning go routines")
	go sayHello()
	go sayHi()

	time.Sleep(1000 * time.Millisecond) // stimulating some work

}
