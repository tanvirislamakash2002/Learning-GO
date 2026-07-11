package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("POST /createUser", createUserHandler)
	mux.HandleFunc("GET /users", getUsersHandler)

	fmt.Println("Server is running at port 5000")
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println("server error", mux)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go server!")
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "server is up and healthy!")
}
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	fmt.Fprintln(w, "Method not allowed")
	// 	return
	// }

	// fmt.Fprintln(w, "User created")

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request body")
		return
	}

	fmt.Println(newUser)

	newUser.Id = len(users) + 1
	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// users, _ := json.Marshal(users)
	// w.Write(users)

	json.NewEncoder(w).Encode(users)
	// encoder.Encode(users)
}
