package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,banana,orange"

	parts := strings.Split(data, ",")
	fmt.Println("Parts value is", parts)

	// checking some other features

	str := "one two three two two"
	count := strings.Count(str, "two")
	fmt.Println("count of two is", count)

	str1 := "   hello owrld    "

	trim := strings.TrimSpace(str1)
	fmt.Println("the value of trim is", trim)

	str2 := "pribce"
	str3 := "kamboj"

	result := strings.Join([]string{str2, str3}, " ")
	fmt.Println("results strings is ", result)

}
