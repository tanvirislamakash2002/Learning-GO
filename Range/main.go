package main

import "fmt"

func main() {
	// myMap := map[string]string{
	// 	"Name":    "Akash",
	// 	"Success": "ok",
	// }

	// for key, value := range myMap {
	// 	fmt.Println(key, value)
	// }
	// for _, value := range myMap { // use ( _ ) if you don't want to use key
	// 	fmt.Println(value)
	// }
	// myArr := [3]string{
	// 	"green",
	// 	"red",
	// 	"black",
	// }

	// colors := []string{
	// 	"green",
	// 	"red",
	// 	"black",
	// }

	// for i, color := range colors {
	// 	fmt.Println(i, color)
	// }

	name := "WebDeveloper"
	var byteSlice = []byte(name)
	// for i, value := range name {
	// 	fmt.Println(i, value)
	// }
	fmt.Println(byteSlice)
}
