package main

import (
	"fmt"
	"sort"
	"strings"
)

func stringsPackage() {
	word := "Hello World"
	fmt.Println(strings.ToUpper(word))
	fmt.Println(strings.ToLower(word))
	fmt.Println(strings.Contains(word, "Hello"))
	fmt.Println(strings.ReplaceAll(word, "Hello", "Hi"))
	fmt.Println(strings.Split(word, " "))
	fmt.Println(strings.Index(word, " "))
}

func sortPackage() {
	numbers := []int{5, 2, 8, 1}

	sort.Ints(numbers)
	fmt.Println(numbers)

	names := []string{"banana", "apple", "orange"}

	sort.Strings(names)
	fmt.Println(names)

	index := sort.SearchInts(numbers, 5)
	fmt.Println(index)
}
