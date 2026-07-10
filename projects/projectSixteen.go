package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputOne() {

	var name string

	fmt.Print("Enter your name: ")

	fmt.Scan(&name)

	fmt.Println("Hello", name)
}

func inputTwo() {

	var age int

	fmt.Print("Enter your age: ")

	fmt.Scan(&age)

	fmt.Println("Your age is", age)
}

func inputThree() {

	var name string
	var age int

	fmt.Print("Enter your name and age: ")

	fmt.Scan(&name, &age)

	fmt.Println(name)
	fmt.Println(age)
}

func inputFour() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a message: ")

	message, _ := reader.ReadString('\n')

	fmt.Println(message)
}
