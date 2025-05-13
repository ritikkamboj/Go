package main

import "fmt"

func Simple() {
	fmt.Println("Jai baabe ki")
}

// func add(a, b int) int {
// 	return a + b
// }

func add(a, b int) (result int) {
	result = a + b

	return
}
func multi(a, b int) int {

	return a * b
}

func main() {
	fmt.Println("we are learning functions in golang ")
	Simple()

	ans := add(3, 4)
	fmt.Println("the vakue of additon is ", ans)

	mul := multi(3, 4)
	fmt.Println("mutiply value is ", mul)
}
