package main

import "fmt"

type Animal interface {
	speak()
}

type Dog struct{}
type Cat struct{}
type Human struct {
	name string
}

// dog
func (d Dog) speak() {
	fmt.Println("Woof!!!")
}

// cat
func (c Cat) speak() {
	fmt.Println("Meow!!!")
}
func (h Human) speak() {
	fmt.Println("My name is!!!", h.name)
}

func makeSound(a Animal) {
	a.speak()
}

func main() {
	dexter := Dog{}
	billi := Cat{}
	akash := Human{name: "Akash"}
	makeSound(dexter)
	makeSound(billi)
	makeSound(akash)
}
