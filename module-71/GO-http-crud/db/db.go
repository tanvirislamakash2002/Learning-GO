package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Db *pgx.Conn

func ConnectDb() {
	var err error
	connStr := os.Getenv("DB_STRING")

	Db, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
		// fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		// os.Exit(1)
	}

	fmt.Println("Database connected successfully")
}
