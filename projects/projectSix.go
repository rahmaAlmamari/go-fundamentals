package main

import "fmt"

func loop() {

	fmt.Println("Traditional for Loop")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Print("for as a while Loop")
	count := 0

	for count < 5 {
		fmt.Println(count)
		count++
	}

	fmt.Println("for range Loop")
	fmt.Println("Loop through a slice")

	numbers := []int{10, 20, 30}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	fmt.Println("Ignore the index")
	for _, value := range numbers {
		fmt.Println(value)
	}

	fmt.Println("Loop through a string")
	for index, char := range "Go" {
		fmt.Println(index, string(char))
	}

	fmt.Println("Loop through a map")
	ages := map[string]int{
		"Ali":  25,
		"Sara": 22,
	}

	for name, age := range ages {
		fmt.Println(name, age)
	}

	fmt.Println("continue and break")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}

		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}

		fmt.Println(i)
	}
}
