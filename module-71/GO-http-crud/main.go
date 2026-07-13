package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
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
	mux.HandleFunc("GET /users/{id}", getSingleUsersHandler)
	mux.HandleFunc("PUT /users/{id}", updateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

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
func getSingleUsersHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	// fmt.Printf("the value or id is %v and the type of the id is %T", idParam, idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
	}

	for _, user := range users {
		if user.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "User not found")
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	// fmt.Printf("the value or id is %v and the type of the id is %T", idParam, idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
		return
	}

	var updatedUser User

	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid request body")
		return
	}

	for idx, user := range users {
		if user.Id == id {
			updatedUser.Id = id
			users[idx] = updatedUser

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "User not found")
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	// fmt.Printf("the value or id is %v and the type of the id is %T", idParam, idParam)

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
		return
	}

	for idx, user := range users {
		if user.Id == id {
			// users = append(users[:idx], users[idx+1:]...)

			users = slices.Delete(users, idx, idx+1)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "User not found")
}
