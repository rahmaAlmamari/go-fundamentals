package main

import "fmt"

func fmtPrint() {
	fmt.Print("Hi every one,")
	fmt.Print(" It is nice to see you all")
}

func fmtPrintln() {
	fmt.Println("Hi every one,")
	fmt.Println(" It is nice to see you all")
}

func printVariable() {
	age := 35
	name := "Ali"

	fmt.Print(name, " is ", age, " years old\n")
}

func fmtPrintF() {
	age := 35
	name := "Ali"
	examMark := 25.64
	finalMark := 77.127

	fmt.Printf("Printf with v -> my age is %v and my name is %v\n", age, name)
	fmt.Printf("Printf with q -> my age is %q and my name is %q\n", age, name)
	fmt.Printf("Printf with T -> the varible store age is in %T data type\n", age)
	fmt.Printf("Printf with f -> your exam mark is %f, and your final mark is %00.2f\n", examMark, finalMark)

}

func fmtSprintf() {
	name := "Ali"
	age := 34
	var str string = fmt.Sprintf("my name is %v, and my age is %v", name, age)
	fmt.Println("the saved string is: ", str)
}
