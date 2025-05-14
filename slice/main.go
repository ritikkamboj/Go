package main

import "fmt"

func main() {

	//slice
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("the use of slice is ", numbers)

	// type of this sloce
	fmt.Printf("type of the array is %T\n", numbers)
	fmt.Println("lenght is ", len(numbers))

	// appending in the array

	numbers = append(numbers, 3, 10, 34, 6, 6)
	fmt.Println("array after appending is ", numbers)
	fmt.Println("lenght is ", len(numbers))

	// seeing the lenght and capacity and make function

	items := []int{1, 2, 3}
	fmt.Println("this is the length of the array us ", len(items))
	fmt.Println("the capacity of the array us ", cap(items))

	// with make function we can chnage the value of capacity

	systems := make([]int, 3, 5)
	fmt.Println("seeing thr length vs capacity ", systems)
	systems = append(systems, 6)
	fmt.Println("seeing thr length vs capacity ", systems)
	fmt.Println("this is the length of the array us ", len(systems))
	fmt.Println("the capacity of the array us ", cap(systems))
	systems = append(systems, 7, 8)
	fmt.Println("seeing thr length vs capacity ", systems)
	fmt.Println("this is the length of the array us ", len(systems))
	fmt.Println("the capacity of the array us ", cap(systems))

	// from above we see that whenw the oversized the array with more values
	// then its capacity gets doubled

	// to initialize empty way we have two ways

	stock := []string{}
	fmt.Println("seeing thr length vs capacity ", stock)

	// other way is by using make f unction

	stock1 := make([]string, 0)
	fmt.Println("seeing thr length vs capacity ", stock1)
	fmt.Println("this is the length of the array us ", len(stock1))
	fmt.Println("the capacity of the array us ", cap(stock1))

}
