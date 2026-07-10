package main

import (
	"fmt"
	"log"
)

func doSomething() {
	defer func() {
		fmt.Println("Differ function ran!")
		r := recover()
		if r != nil {
			fmt.Println("Recovered from", r)
		}
	}()
	panic("something really bad happen")
}

func doAnotherThing() {
	defer func() {
		fmt.Println("Differ function ran!")
	}()
	log.Fatal("Something Very Big Happened")
}
func main() {
	// doSomething()
	doAnotherThing()

	fmt.Println("main complete normally")
}
