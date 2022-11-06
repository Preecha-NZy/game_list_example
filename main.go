package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DbUsername = "preechac"
	DbPassword = "0836161840NZy"
	DbUrl      = "postgres://" + DbUsername + ":" + DbPassword + "@127.0.0.1:5432/gle?sslmode=disable"
)

func main() {
	_, err := sql.Open("postgres", DbUrl)

	if err != nil {
		fmt.Println("databasee connect error", err)
	}

	fmt.Println("Hello, World!")
	// router := makeRoute()
	// http.Handle("/", router)
	// http.ListenAndServe(":8000", router)
}
