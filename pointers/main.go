package main

import "fmt"

func modifyRef(ptr *int) {
	*ptr = *ptr + 20
}

func main() {

	fmt.Println("studying pointers")

	var name int
	name = 2

	var ptr *int
	ptr = &name

	fmt.Println("Value of name is  ", name)
	fmt.Println("value of ptr is ", ptr)

	// other way to define pointers

	name1 := 3
	ptr1 := &name1

	fmt.Println("ptr has value", name1)
	fmt.Println(" pointer has value ", ptr1)
	fmt.Println(" pointer has value ", *ptr1)

	var ptr3 *int

	if ptr3 == nil {
		fmt.Println("Value is nil")

	}

	// below is more importamt

	value := 4

	modifyRef(&value)

	fmt.Println("Valkue is ", value)
}
