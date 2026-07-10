package main

import "fmt"

var message1 = "Hello from package scope"

func showMessage() {
	fmt.Println(message1)
}

var shared = "I am shared across files"

func printShared() {
	fmt.Println(shared)
}
