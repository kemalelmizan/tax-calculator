package controller

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

// ConnectionURL ...
func ConnectionURL(username, password, dbname, host string, port int) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", username, password, host, port, dbname)
}

// Migrate ...
func Migrate() {
	db, err := sql.Open("postgres", ConnectionURL("postgres", "dummypassword", "tax_calculator", "localhost", 5432))
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///../migrations",
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	m.Steps(2)
	fmt.Println("Migration successful.")
}
