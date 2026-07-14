// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"strconv"

// 	"github.com/jackc/pgx/v5"
// 	"github.com/joho/godotenv"
// )

// type User struct {
// 	Id    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Email string `json:"email"`
// }

// var users = []User{
// 	{
// 		Id:    1,
// 		Name:  "Kamal",
// 		Age:   25,
// 		Email: "kamal@gmail.com",
// 	},
// 	{
// 		Id:    2,
// 		Name:  "jamal",
// 		Age:   25,
// 		Email: "jamal@gmail.com",
// 	},
// }

// var db *pgx.Conn

// func connectDb() {
// 	var err error
// 	connStr := os.Getenv("DB_STRING")

// 	db, err = pgx.Connect(context.Background(), connStr)
// 	if err != nil {
// 		panic(err)
// 		// fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		// os.Exit(1)
// 	}

// 	fmt.Println("Database connected successfully")
// }

// func main() {

// 	var err error

// 	err = godotenv.Load()

// 	if err != nil {
// 		panic("env file not found")
// 	}

// 	connectDb()

// 	defer db.Close(context.Background())
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("GET /", rootHandler)
// 	mux.HandleFunc("GET /health", healthHandler)
// 	mux.HandleFunc("POST /createUser", createUserHandler)
// 	mux.HandleFunc("GET /users", getUsersHandler)
// 	mux.HandleFunc("GET /users/{id}", getSingleUsersHandler)
// 	mux.HandleFunc("PUT /users/{id}", updateUserHandler)
// 	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

// 	fmt.Println("Server is running at port 5000")
// 	err = http.ListenAndServe(":5000", mux)
// 	if err != nil {
// 		fmt.Println("server error", mux)
// 	}
// }

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Welcome to go server!")
// }
// func healthHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "server is up and healthy!")
// }
// func createUserHandler(w http.ResponseWriter, r *http.Request) {
// 	// if r.Method != "POST" {
// 	// 	w.WriteHeader(http.StatusMethodNotAllowed)
// 	// 	fmt.Fprintln(w, "Method not allowed")
// 	// 	return
// 	// }

// 	// fmt.Fprintln(w, "User created")

// 	var newUser User

// 	err := json.NewDecoder(r.Body).Decode(&newUser)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintln(w, "Invalid request body")
// 		return
// 	}

// 	fmt.Println(newUser)

// 	// newUser.Id = len(users) + 1
// 	// users = append(users, newUser)

// 	query := `
// 	insert into users (username, age, email)
// 	values ($1, $2, $3)
// 	returning id
// 	`

// 	err = db.QueryRow(context.Background(), query, newUser.Name, newUser.Age, newUser.Email).Scan(&newUser.Id)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintln(w, "Could not create user")
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(newUser)

// }
// func getUsersHandler(w http.ResponseWriter, r *http.Request) {

// 	query := `select id, username, age, email from users`
// 	rows, err := db.Query(context.Background(), query)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintln(w, "Could not get users")
// 		return
// 	}

// 	defer rows.Close()

// 	var users []User

// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			fmt.Fprintln(w, "Could not scan user")
// 			return
// 		}
// 		users = append(users, user)

// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintln(w, "Could not read user")
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	// users, _ := json.Marshal(users)
// 	// w.Write(users)

// 	json.NewEncoder(w).Encode(users)
// 	// encoder.Encode(users)
// }
// func getSingleUsersHandler(w http.ResponseWriter, r *http.Request) {
// 	idParam := r.PathValue("id")
// 	// fmt.Printf("the value or id is %v and the type of the id is %T", idParam, idParam)

// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintln(w, "Invalid user id")
// 		return
// 	}

// 	// for _, user := range users {
// 	// 	if user.Id == id {
// 	// 		w.Header().Set("Content-Type", "application/json")
// 	// 		json.NewEncoder(w).Encode(user)
// 	// 	}
// 	// }

// 	var user User
// 	query := `SELECT id, username, age, email FROM users WHERE id = $1`

// 	err = db.QueryRow(context.Background(), query, id).Scan(&user.Id, &user.Name, &user.Age, &user.Email)

// 	if err == pgx.ErrNoRows {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintln(w, "User not found")
// 		return
// 	}

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintln(w, "could not get user")
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(user)

// }

// func updateUserHandler(w http.ResponseWriter, r *http.Request) {
// 	idParam := r.PathValue("id")
// 	// fmt.Printf("the value or id is %v and the type of the id is %T", idParam, idParam)

// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintln(w, "Invalid user id")
// 		return
// 	}

// 	var updatedUser User

// 	err = json.NewDecoder(r.Body).Decode(&updatedUser)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintln(w, "Invalid request body")
// 		return
// 	}

// 	query := `
// 	update users
// 	set username = $1, age = $2, email = $3
// 	where id = $4
// 	returning id, username, age, email
// 	`
// 	err = db.QueryRow(context.Background(), query, updatedUser.Name, updatedUser.Age, updatedUser.Email, id).Scan(&updatedUser.Id, &updatedUser.Name, &updatedUser.Age, &updatedUser.Email)

// 	if err == pgx.ErrNoRows {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintln(w, "user not found")
// 		return
// 	}

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintln(w, "1Could not  update user")
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(updatedUser)
// }

// func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
// 	idParam := r.PathValue("id")
// 	// fmt.Printf("the value or id is %v and the type of the id is %T", idParam, idParam)

// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		fmt.Fprintln(w, "Invalid user id")
// 		return
// 	}

// 	// for idx, user := range users {
// 	// 	if user.Id == id {
// 	// 		// users = append(users[:idx], users[idx+1:]...)

// 	// 		users = slices.Delete(users, idx, idx+1)
// 	// 		w.WriteHeader(http.StatusNoContent)
// 	// 		return
// 	// 	}
// 	// }

// 	query := `delete from users where id = $1`
// 	cmdTag, err := db.Exec(context.Background(), query, id)

// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		fmt.Fprintln(w, "Could not delete user")
// 		return
// 	}

// 	if cmdTag.RowsAffected() != 1 {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintln(w, "user not found")
// 		return
// 	}
// 	w.WriteHeader(http.StatusNoContent)
// 	fmt.Fprintln(w, "User deleted successfully")
// }
