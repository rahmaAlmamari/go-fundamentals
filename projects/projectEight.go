package main

import (
	"fmt"
	"strings"
)

func show(n string) {
	fmt.Printf("Hi, I am %v", n)
}

var Names = []string{"Rahma", "Fahad", "Ali"}

func sayHi(n string) {
	fmt.Printf("Hi, I am %v \n", n)
}
func message(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func area(r float64) float64 {
	return r * 2
}

func getTnitials(n string) (string, string) {
	upper := strings.ToUpper(n)
	words := strings.Split(upper, " ")

	var initials []string
	for _, v := range words {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}
