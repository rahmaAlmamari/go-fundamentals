package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func typeOne() {

	user := User{
		Name: "Rahma",
		Age:  24,
	}

	//fmt.Println(user)

	// fmt.Println(user.Name)
	// fmt.Println(user.Age)

	user.Age = 25
	fmt.Println(user.Age)
}

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func typeTwo() {

	employee := Employee{
		ID:         1,
		Name:       "Ali",
		Department: "IT",
		Salary:     900,
	}

	fmt.Println(employee.Name)
}

type typeUser struct {
	Name string
}

func (u typeUser) SayHello() {
	fmt.Println("Hello", u.Name)
}

func typeThree() {

	user := typeUser{
		Name: "Rahma",
	}

	user.SayHello()
}

func updateName(user *typeUser) {
	user.Name = "Ahmed"
}

func typeFour() {

	user := typeUser{
		Name: "Rahma",
	}

	updateName(&user)

	fmt.Println(user.Name)
}
