package main

import "fmt"

func pointerOne() {
	name := "Rahma"

	fmt.Println(name)
	fmt.Println(&name)
}

func pointerTwo() {
	age := 24

	ptr := &age

	fmt.Println(ptr)
	fmt.Println(*ptr)
}

func pointerThree() {
	number := 100

	ptr := &number

	fmt.Println("Value:", number)
	fmt.Println("Address:", ptr)
	fmt.Println("Dereferenced:", *ptr)
}

func pointerFour() {
	number := 50

	ptr := &number

	*ptr = 200

	fmt.Println(number)
}

func increasePointer(num *int) {
	*num = *num + 10
}

func pointerFive() {
	number := 20

	increasePointer(&number)

	fmt.Println(number)
}
