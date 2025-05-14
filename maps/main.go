package main

import "fmt"

func main() {

	studentGrades := make(map[string]int)

	studentGrades["Ash"] = 70
	studentGrades["Ash1"] = 701
	studentGrades["Ash2"] = 702
	studentGrades["Ash3"] = 703

	fmt.Println("map is ", studentGrades)
	fmt.Println(" value of one is", studentGrades["Ash"])

	studentGrades["Ash"] = 100

	fmt.Println("new marks is ", studentGrades["Ash"])

	//deleting the values

	// lets delete Ash

	delete(studentGrades, "Ash")
	fmt.Println(" value of one is", studentGrades["Ash"])

	// chekc if key is there

	grade, exists := studentGrades["Ash"]
	fmt.Println("Grades of Ash", grade)
	fmt.Println("is it exits ", exists)

	grade2, exists2 := studentGrades["Ash1"]
	fmt.Println("Grades of Ash1", grade2)
	fmt.Println("is it exits ", exists2)

	for index, value := range studentGrades {
		fmt.Printf("key is %s and marks is %d\n", index, value)
	}

	//earier we declare then we define and now we have to define as we are declaring

	students := map[string]int{
		"Alice": 100,
		"Bob":   89,
		"Megha": 23,
	}

	for index, value := range students {
		fmt.Printf("the index is %s and value is %d\n", index, value)
	}

}
