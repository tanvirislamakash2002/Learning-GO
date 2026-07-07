package main

import "fmt"

func main() {
	var numbers [6]int
	numbers[1] = 63
	numbers[3] = 22
	fmt.Println(numbers)
	fmt.Println(numbers[3])

	fmt.Println(len(numbers))

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
