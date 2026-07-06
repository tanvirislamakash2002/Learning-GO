package main

import "fmt"

// Variadic function
func add(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func greet(prefix string, mps ...string) {
	for _, mp := range mps {
		fmt.Println(prefix, mp)
	}
}

func main() {
	sum := add(5, 7, 12)
	fmt.Println(sum)
	mps := []string{"Jamal", "Kamal", "Khairul", "mizan"}
	greet("Welcome", mps...) //variadic argument
}

// flexible amount of argument
// must be the last parameter
// internally slice
