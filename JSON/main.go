package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"personName"`
	Age  int    `json:"-"`
	City string `json:"city"`
}

func main() {
	p := person{
		Name: "Akash",
		Age:  25,
		City: "Dhaka",
	}
	rawJson, err := json.Marshal(p)

	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(string(rawJson))
}
