package main

import "fmt"

func main() {
	// var numbers = [6]int{10, 20, 30} //partial array initialization
	// var orders = [6]int{10, 20, 30, 40, 50, 60}

	// slice1 := orders[3:4] // [start:end(excluded)] len, cap, pointer
	// fmt.Println(slice1)

	// slice1 = append(slice1, 300)
	// slice1 = append(slice1, 400)
	// slice1 = append(slice1, 500)
	// slice1 = append(slice1, 600)
	// fmt.Println("the length of the slice is ", len(slice1))
	// fmt.Println("the cap of the slice is ", cap(slice1))
	// fmt.Println(slice1)
	// fmt.Println(orders)

	var slice1 = []int{1, 2, 3}
	slice1 = append(slice1, 200)
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
}
