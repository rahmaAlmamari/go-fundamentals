package main

import "fmt"

func boolExample() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)
}

func ifExample() {
	age := 45

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is less than 50")
	}
}
