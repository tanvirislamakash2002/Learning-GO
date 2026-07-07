package main

import "fmt"

// func init() {
// 	fmt.Println("Initializing db...")
// }

func process(sayHello func()) {
	sayHello()
}

func calculate(a int, b int, operation func(x int, y int) int) int {
	return operation(a, b)
}
func main() {
	// color.Cyan("Prints text in cyan")
	// color.Blue("Prints text in Blue")
	// color.Red("Prints text in red")

	// greet := func() {
	// 	fmt.Println("Hello there")
	// }
	// process(greet)

	add := func(n1 int, n2 int) int {
		return n1 + n2
	}
	multiply := func(n1 int, n2 int) int {
		return n1 * n2
	}
	fmt.Println(calculate(10, 20, add))
	fmt.Println(calculate(10, 20, multiply))
}

// init and main function can run without call
