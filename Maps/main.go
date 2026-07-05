package main

import "fmt"

type user struct {
	name       string
	age        int
	isLoggedIn bool
}

func main() {
	// myMap := make(map[string]int)

	// myMap["user1Score"] = 5
	// myMap["user2Score"] = 7

	// fmt.Println(myMap)
	// fmt.Println(myMap["user1Score"])

	// myMap := map[string]string{
	// 	"name":    "Akash",
	// 	"Success": "ok",
	// }
	// delete(myMap, "name")

	myMap := map[string]user{
		"data": user{
			name:       "mizan",
			age:        25,
			isLoggedIn: false,
		},
	}
	fmt.Println(myMap)
}
