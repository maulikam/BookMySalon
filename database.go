package main

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "reniladmin"
	dbname   = "bookmysalon"
)




func connect() *sql.DB {
	connStr := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	connStr = fmt.Sprintf(connStr, host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
