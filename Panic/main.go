package main

import "fmt"

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

func main() {
	doSomething()

	fmt.Println("main complete normally")
}
