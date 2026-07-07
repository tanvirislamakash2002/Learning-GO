package main

import "fmt"

func multiplyBy(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
func main() {

	double := multiplyBy(2)
	fmt.Println(double(5))

}
