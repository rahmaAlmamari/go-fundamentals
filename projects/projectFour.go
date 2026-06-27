package main

import "fmt"

func declareArray() {
	var ages = [3]int{20, 25, 30}
	fmt.Println(ages, len(ages))
}

func changeArrayValue() {
	var names = [3]string{"Ali", "Fahad", "Ahmed"}
	fmt.Println(names)

	names[0] = "Rahma"
	fmt.Println(names)
}

func declareSlice() {
	var scores = []int{20, 25, 30}
	fmt.Println(scores)
}

func changeSliceValue() {
	var names = []string{"Ali", "Fahad", "Ahmed"}
	fmt.Println(names)

	names[0] = "Rahma"
	fmt.Println(names)
}

func appendSlice() {
	var scores = []int{20, 25, 30}
	fmt.Println(append(scores, 85))
	fmt.Println(scores)
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))
}

func sliceRanges() {
	var scores = []int{20, 25, 30, 10, 85}
	marks := scores[1:3]
	fmt.Println(marks)
	marksTwo := scores[2:]
	fmt.Println(marksTwo)
	marksThree := scores[:3]
	fmt.Println(marksThree)
}
