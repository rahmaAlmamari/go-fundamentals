package main

import (
	"fmt"
	"strconv"
)

func saveFile() {

	price := "19.99"

	value, err := strconv.ParseFloat(price, 64)

	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	fmt.Println(value)
}

func saveFileTwo() {

	value, err := strconv.ParseFloat("hello", 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(value)
}
