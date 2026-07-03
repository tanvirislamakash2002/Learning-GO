package main

import "fmt"

func makeCoffee(kind string) {
	fmt.Printf("Making %s Coffee...\n", kind)
}
func makeCoffeeReturn(kind string) string {
	coffee := fmt.Sprintf("%s Coffee", kind)
	return coffee
}
func makeCoffeeReturnMV(kind string) (string, int) {
	price := 25
	coffee := fmt.Sprintf("%s Coffee", kind)
	return coffee, price
}
func makeCoffeeReturnMVN(kind string) (coffee string, price int) { //named return value
	price = 25
	coffee = fmt.Sprintf("%s Coffee", kind)
	return
}
func main() {

	// var name string
	// name = "mokless"

	// var designation string = "Developer"

	// var id = 12122
	// age := 22

	// var (
	// 	address  string = "Rajshahi"
	// 	postCode int    = 5000
	// )

	// var x, y int

	// x = 2
	// y = 5

	// var p, q int = 7, 1

	// const pi = 3.1416

	// formattedString := fmt.Sprintf("2nd My name Is %s and I am %d years old. I am a %s . I know the value of pi which is %.3f", name, age, designation, pi)

	// fmt.Println(pi)
	// fmt.Println(x, y)
	// fmt.Println(p, q)

	// fmt.Println(name, designation, id, age)
	// fmt.Println(address, postCode)

	// fmt.Printf("My name Is %s and I am %d years old. I am a %s . I know the value of pi which is %.3f", name, age, designation, pi)

	// fmt.Println(formattedString)

	// function ------------------------------

	// makeCoffee("cold")
	// makeCoffee("black")

	// myCoffee := makeCoffeeReturn("black&white")
	// fmt.Println(myCoffee)

	// myChooseCoffee, myBill := makeCoffeeReturnMV("brown")
	// fmt.Printf("I am having a %d$ %s", myBill, myChooseCoffee)

	// coffeeBanao := func() {
	// 	fmt.Printf("making the best coffee")
	// }
	// coffeeBanao()

	// func(coffeeType string) {
	// 	fmt.Printf("making the best coffee I Ever test, which is %s", coffeeType)
	// }("Latte")

}
