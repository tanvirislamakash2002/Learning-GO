package main

import (
	"context"
	"fmt"
	"go-http-crud/db"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	var err error

	err = godotenv.Load()

	if err != nil {
		panic("env file not found")
	}
	db.ConnectDb()

	defer db.Db.Close(context.Background())
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("POST /createUser", createUserHandler)
	mux.HandleFunc("GET /users", getUsersHandler)
	mux.HandleFunc("GET /users/{id}", getSingleUsersHandler)
	mux.HandleFunc("PUT /users/{id}", updateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

	fmt.Println("Server is running at port 5000")
	err = http.ListenAndServe(":5000", mux)
	if err != nil {
		fmt.Println("server error", mux)
	}
}
