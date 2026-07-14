package main

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var users = []User{
	{
		Id:    1,
		Name:  "Kamal",
		Age:   25,
		Email: "kamal@gmail.com",
	},
	{
		Id:    2,
		Name:  "jamal",
		Age:   25,
		Email: "jamal@gmail.com",
	},
}
