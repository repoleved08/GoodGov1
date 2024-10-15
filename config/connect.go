package config

import (
	"database/sql"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
// setup mysql dsn
dsn := fmt.Sprintf
