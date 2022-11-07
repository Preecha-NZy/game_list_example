package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const (
	DbUsername = "preechac"
	DbPassword = "0836161840NZy"
	DbUrl      = "postgres://" + DbUsername + ":" + DbPassword + "@127.0.0.1:5432/gle?sslmode=disable"
)

func connect() (*sql.DB, error) {
	return sql.Open("postgres", DbUrl)
}
