package main

import (
	"fmt"
	"strconv"
)

func parseFloat() {

	price := "19.99"

	value, err := strconv.ParseFloat(price, 64)

	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	fmt.Println(value)
}

func parseFloatTwo() {

	value, err := strconv.ParseFloat("hello", 64)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(value)
}
