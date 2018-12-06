package model

import (
	"os"
	"testing"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func Test_dbEnv(t *testing.T) {
	tests := []struct {
		name              string
		wantUsername      string
		wantPassword      string
		wantDbname        string
		wantHost          string
		wantPort          int
		wantConnectionURL string
	}{
		{
			name:              "Happy path",
			wantUsername:      "a",
			wantPassword:      "b",
			wantDbname:        "c",
			wantHost:          "d",
			wantPort:          1,
			wantConnectionURL: "postgres://a:b@d:1/c?sslmode=disable",
		},
	}
	for _, tt := range tests {
		os.Setenv("db_host", "d")
		defer os.Unsetenv("db_host")

		os.Setenv("db_port", "1")
		defer os.Unsetenv("db_port")

		os.Setenv("db_username", "a")
		defer os.Unsetenv("db_username")

		os.Setenv("db_password", "b")
		defer os.Unsetenv("db_password")

		os.Setenv("db_name", "c")
		defer os.Unsetenv("db_name")

		t.Run(tt.name, func(t *testing.T) {
			gotUsername, gotPassword, gotDbname, gotHost, gotPort, gotConnectionURL := dbEnv()
			if gotUsername != tt.wantUsername {
				t.Errorf("dbEnv() gotUsername = %v, want %v", gotUsername, tt.wantUsername)
			}
			if gotPassword != tt.wantPassword {
				t.Errorf("dbEnv() gotPassword = %v, want %v", gotPassword, tt.wantPassword)
			}
			if gotDbname != tt.wantDbname {
				t.Errorf("dbEnv() gotDbname = %v, want %v", gotDbname, tt.wantDbname)
			}
			if gotHost != tt.wantHost {
				t.Errorf("dbEnv() gotHost = %v, want %v", gotHost, tt.wantHost)
			}
			if gotPort != tt.wantPort {
				t.Errorf("dbEnv() gotPort = %v, want %v", gotPort, tt.wantPort)
			}
			if gotConnectionURL != tt.wantConnectionURL {
				t.Errorf("dbEnv() gotConnectionURL = %v, want %v", gotConnectionURL, tt.wantConnectionURL)
			}
		})
	}
}

func TestInitDB(t *testing.T) {
	_, err := InitDB()
	require.NoError(t, err)
}

func TestMigrate(t *testing.T) {
	os.Setenv("migrations_dir", "/dev/null")
	defer os.Unsetenv("migrations_dir")
	require.Panics(t, func() {
		Migrate()
	})
}
