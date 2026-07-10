package main

import (
	"fmt"
)

type CustomError struct {
	message string
	code    int
}

func (cu *CustomError) Error() string {
	return cu.message
}
func login(password string) error {
	if password != "1234" {
		return &CustomError{
			message: "Password do not match",
			code:    401,
		}
		// return errors.New("password do not match")
	}
	return nil
}
func main() {
	err := login("4435")
	if err != nil {
		// customErr, ok := err.(*CustomError)

		// if !ok {
		// 	fmt.Println(customErr)
		// } else {
		// 	fmt.Println(customErr.message)
		// 	fmt.Println(customErr.code)
		// }

		if CustomError, ok := err.(*CustomError); ok {
			fmt.Println(CustomError.code)
		}
		// fmt.Println("Error", err, "Code", err.code)
	}

	users := map[int]string{
		1: "Tanvir",
		2: "Islam",
		3: "Akash",
	}
	name, ok1 := users[3]
	name2, ok := users[4]
	fmt.Println(name, ok1)
	fmt.Println("name: ", name2, ok)
	fmt.Println("main ends")
}
