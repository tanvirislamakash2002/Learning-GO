package main

import "fmt"

// func change(x *int) {
// 	*x = 400
// }
func main() {
	// a := 42
	// p := &a // 0x10ff46360a0 --> pointer (&)
	// a = 100

	// fmt.Println("a:", a)
	// fmt.Println("p:", p)
	// fmt.Println("p:", *p) // dereferencing (*)
	// fmt.Println(&a)

	// y := 20

	// change(&y)

	// fmt.Println(y)

	bigArray := [5]int{10, 20, 30, 40, 50}

	// modifyWithoutPointer(bigArray)
	// fmt.Println("After without pointer:", bigArray)

	modifyWithPointer(&bigArray)
	fmt.Println("After with pointer:", bigArray)
}

func modifyWithoutPointer(arr [5]int) {
	arr[0] = 999
	fmt.Println("Inside without pointer:", arr)
}

func modifyWithPointer(arr *[5]int) {
	arr[0] = 444
	fmt.Println("Inside without pointer:", *arr)
}
