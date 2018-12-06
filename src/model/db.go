package model

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate"

	// postgres, file and pq import for migration
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func dbEnv() (username, password, dbname, host string, port int, connectionURL string) {
	host = os.Getenv("db_host")
	port, _ = strconv.Atoi(os.Getenv("db_port"))
	username = os.Getenv("db_username")
	password = os.Getenv("db_password")
	dbname = os.Getenv("db_name")
	connectionURL = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", username, password, host, port, dbname)

	return username, password, dbname, host, port, connectionURL
}

// InitDB ...
func InitDB() (*sql.DB, error) {
	username, password, dbname, host, port, _ := dbEnv()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db, err
}

// Migrate ...
func Migrate() {

	_, _, _, _, _, connectionURL := dbEnv()
	m, err := migrate.New(os.Getenv("migrations_dir"), connectionURL)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Migration successful.")
}
