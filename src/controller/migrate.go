package controller

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

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
	port, _ := strconv.Atoi(os.Getenv("db_port"))
	db, err := sql.Open("postgres", ConnectionURL(os.Getenv("db_username"), os.Getenv("db_password"), os.Getenv("db_name"), os.Getenv("db_host"), port))
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		os.Getenv("migrations_dir"),
		"postgres", driver)
	if err != nil {
		panic(err)
	}
	m.Steps(2)
	fmt.Println("Migration successful.")
}
