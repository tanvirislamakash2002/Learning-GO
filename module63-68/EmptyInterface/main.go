package main

import "fmt"

// any = interface{}
// func Print(data interface{}) {
// 	fmt.Println(data)
// }
// func Print(data any) {
// 	fmt.Println(data)
// }

// type assertion
func Process(data any) {
	strData, ok := data.(string)
	if ok {
		fmt.Println(strData)
		fmt.Println(len(strData))
	}
	intData, ok := data.(int)
	if ok {
		fmt.Println(intData + 100)
	}

}
func main() {
	// var data interface{}
	// data = "Akash"
	// fmt.Println(data)

	// Print([]int{3, 4, 6, 1})
	// Print("Tanvir Islam")

	Process(22)
	Process("Alpha")
}
