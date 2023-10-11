package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "reniladmin"
	dbname   = "bookmysalon"
)

// Connect establishes a connection to the database and returns it.
func Connect() (*sql.DB, error) {
	connStr := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	connStr = fmt.Sprintf(connStr, host, port, user, password, dbname)
	return sql.Open("postgres", connStr)
}
