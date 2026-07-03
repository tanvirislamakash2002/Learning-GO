package main

import "fmt"

func makeCoffee(coffeeNo int) {
	fmt.Println("making coffee no", coffeeNo)
}
func main() {
	for i := 1; i <= 5; i++ {
		makeCoffee(i)
		fmt.Println(i)
	}
}
