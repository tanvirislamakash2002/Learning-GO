package main

import "fmt"

func calculate(a int, b int, operation func(x int, y int) int) int {
	return operation(a, b)
}

func multiplyBy(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
func main() {

	// add := func(n1 int, n2 int) int {
	// 	return n1 + n2
	// }

	// fmt.Println(calculate(10, 20, add))

	double := multiplyBy(2)
	fmt.Println(double(5))

}
