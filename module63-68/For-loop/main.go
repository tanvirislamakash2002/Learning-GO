package main

import "fmt"

func makeCoffee(coffeeNo int) {
	fmt.Println("making coffee no1", coffeeNo)
}
func main() {
	// for i := 1; i <= 5; i++ {
	// 	makeCoffee(i)
	// 	fmt.Println(i)
	// }

	// i := 1

	// for i <= 5 {
	// 	makeCoffee(i)
	// 	i++
	// }

	// infinite loop
	// for {
	// 	makeCoffee(2)
	// }

	// break (break the loop) and continue (skip the iteration)
	// for i := 1; i <= 10; i++ {
	// 	makeCoffee(i)
	// 	if i == 6 {
	// 		break
	// 	}
	// }
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		makeCoffee(i)
	}
}
