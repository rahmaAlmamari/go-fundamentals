package main

import "fmt"

func increase(num int) {
	num = num + 10
	fmt.Println("Inside function:", num)
}

func passOne() {
	number := 20

	increase(number)

	fmt.Println("Outside function:", number)
}

func changeName(name string) {
	name = "Ahmed"
	fmt.Println("Inside function:", name)
}

func passTwo() {
	username := "Rahma"

	changeName(username)

	fmt.Println("Outside function:", username)
}

func enable(status bool) {
	status = true
	fmt.Println("Inside function:", status)
}

func passThree() {
	isEnabled := false

	enable(isEnabled)

	fmt.Println("Outside function:", isEnabled)
}

func update(name string, age int) {
	name = "Ali"
	age = 30

	fmt.Println(name)
	fmt.Println(age)
}

func passFour() {
	name := "Rahma"
	age := 24

	update(name, age)

	fmt.Println(name)
	fmt.Println(age)
}

func increaseTwo(num int) int {
	return num + 10
}

func passFive() {
	number := 20

	number = increaseTwo(number)

	fmt.Println(number)
}

func change(numbers []int) {
	numbers[0] = 100
}

func passSix() {
	values := []int{1, 2, 3}

	change(values)

	fmt.Println(values)
}

func changeAge(age int) {
	age = 30
}

func passSeven() {
	age := 20

	changeAge(age)

	fmt.Println(age)
}
