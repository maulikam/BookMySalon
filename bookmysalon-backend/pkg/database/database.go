package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func RunMigrations() {
	db, err := Connect()
	if err != nil {
		log.Fatalf("Could not connect to database... %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not start sql migration... %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/database/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Migration failed... %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("An error occurred while migrating... %v", err)
	}
}
