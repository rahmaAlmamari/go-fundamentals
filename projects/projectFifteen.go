package main

import "fmt"

type reUser struct {
	Name string
	Age  int
}

func (u reUser) PrintInfo() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
}

func reOne() {

	user := reUser{
		Name: "Rahma",
		Age:  24,
	}

	// user.PrintInfo()

	// user.ChangeName()

	// fmt.Println(user.Name)

	user.ChangeNameTwo()

	fmt.Println(user.Name)
}

func (u reUser) ChangeName() {
	u.Name = "Ahmed"
}

func (u *reUser) ChangeNameTwo() {
	u.Name = "Ahmed"
}
