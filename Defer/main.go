// package main

// import "fmt"

// // func deferred(result int) {
// // 	fmt.Println("deferred result:", result)
// // }
// func example() int {
// 	result := 10
// 	// defer deferred(result)
// 	defer func() {
// 		result = result + 200
// 		fmt.Println("deferred result:", result)
// 	}()
// 	fmt.Println("I am from example fn:", result)
// 	result = result + 100
// 	return result
// }
// func main() {
// 	// defer fmt.Println("I am deferred")
// 	// fmt.Println("I am from main function")
// 	fmt.Println("return result", example())
// }

// ==> /** result format **/
// I am from example fn: 10
// deferred result: 310
// return result 110

package main

import "fmt"

func example() (result int) {
	result = 10
	// defer deferred(result)
	defer func() {
		result = result + 200
		fmt.Println("deferred result:", result)
	}()
	fmt.Println("I am from example fn:", result)
	result = result + 100
	return
}
func main() {
	// defer fmt.Println("I am deferred")
	// fmt.Println("I am from main function")
	fmt.Println("return result", example())
}

// ==> /** result format **/
// I am from example fn: 10
// deferred result: 310
// return result 310
