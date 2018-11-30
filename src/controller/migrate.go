package controller

import (
	"fmt"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate"

	// postgres, file and pq import for migration
	_ "github.com/golang-migrate/migrate/database/postgres"
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
	m, err := migrate.New(os.Getenv("migrations_dir"), ConnectionURL(os.Getenv("db_username"), os.Getenv("db_password"), os.Getenv("db_name"), os.Getenv("db_host"), port))
	if err != nil {
		fmt.Println(err)
	}
	m.Run()

	fmt.Println("Migration successful.")
}
