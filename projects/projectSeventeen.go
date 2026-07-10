package main

import "fmt"

func switchExa() {

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Almost weekend")
	default:
		fmt.Println("Regular day")
	}
}

func switchExaOne() {

	day := "Saturday"

	switch day {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}
}

func switchExaTwo() {

	age := 20

	switch {
	case age < 18:
		fmt.Println("Minor")
	case age >= 18:
		fmt.Println("Adult")
	}
}

func switchExaThree() {

	number := 1

	switch number {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
	}
}
