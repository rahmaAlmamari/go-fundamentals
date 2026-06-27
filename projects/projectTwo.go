package main

import "fmt"

func printString() {
	var nameOne string = "Rahma"
	fmt.Println(nameOne)

	var nameTwo = "Fahad"
	fmt.Println(nameTwo)

	var nameThree string
	fmt.Println(nameThree)

	nameFour := "Ali"
	fmt.Println(nameFour)

}

func printInts() {

	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)
}

func printFloats() {

	var ageOne float32 = 25.99
	var ageTwo float64 = 30.
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)
}
