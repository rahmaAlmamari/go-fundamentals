package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

func interfaceExa() {
	var s Speaker

	s = Dog{}

	s.Speak()
}

type Cat struct{}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

func makeSound(s Speaker) {
	s.Speak()
}

func interfaceExaTwo() {
	makeSound(Dog{})
	makeSound(Cat{})
}
