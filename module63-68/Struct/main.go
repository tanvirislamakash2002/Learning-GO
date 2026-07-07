package main

import "fmt"

// type additionalInfo struct {
// 	phone   int
// 	address string
// }

// type user struct {
// 	name     string
// 	email    string
// 	metaInfo additionalInfo
// }

// type user struct {
// 	name string
// 	age  int
// 	role string
// }

// type user struct {
// 	name       string
// 	age        int
// 	isLoggedIn bool
// 	greet      func()
// }

type user struct {
	name       string
	age        int
	isLoggedIn bool
}

func main() {
	// john := user{"john", "john@gmail.com"} // positional
	// john := user{name: "john", email: "john@gmail.com"} // key-value based
	// fmt.Println(john)
	// fmt.Printf("%+v", john)
	// fmt.Println(john.email)

	// var user1 user
	// user1.name = "akash"
	// user1.email = "akash@gmail.com"
	// fmt.Println(user1)

	// john := user{
	// 	name:  "john",
	// 	email: "john@gmail.com",
	// 	metaInfo: additionalInfo{
	// 		phone:   01324534437,
	// 		address: "117/A Rajshahi",
	// 	},
	// }

	// fmt.Printf("%+v", john)

	// newUser := func(name string, age int, role string) user {
	// if name == "" {
	// 	fmt.Println("user is required")
	// 	return user{}
	// }
	// 	if age <= 0 {
	// 		age = 18
	// 	}
	// 	return user{
	// 		name: name,
	// 		age:  age,
	// 		role: role,
	// 	}
	// }

	// jamal := newUser("akash", -20, "admin")
	// fmt.Println(jamal)

	// user1 := user{
	// 	name:       "John",
	// 	age:        25,
	// 	isLoggedIn: false,
	// }
	// user1.greet = func() {
	// 	fmt.Println("Hello", user1.name)
	// }
	// user1.greet()

	user1 := user{
		name:       "Akash",
		age:        25,
		isLoggedIn: false,
	}
	// pointerUser1 := &user1
	user1.greet()
	user1.login()

	fmt.Printf("%+v", user1)
}

func (u *user) login() {
	// (*u).isLoggedIn = true
	u.isLoggedIn = true
}

func (u user) greet() {
	fmt.Println("Hello", u.name)
}
