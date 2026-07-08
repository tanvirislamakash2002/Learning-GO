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

	var p2 person

	jsonText := `{"personName":"Akash","city":"Dhaka"}`

	error := json.Unmarshal([]byte(jsonText), &p2)
	if error != nil {
		fmt.Println("Error", err)
	}

	fmt.Println(p2.Age)
	fmt.Println(p2.Name)
	fmt.Println(p2.City)

}
