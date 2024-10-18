package config

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
    dbUser := os.Getenv("DB_USER")
    if dbUser == "" {
        return nil, fmt.Errorf("DB_USER environment variable not set")
    }

    dbPassword := os.Getenv("DB_PASSWORD")
    if dbPassword == "" {
        return nil, fmt.Errorf("DB_PASSWORD environment variable not set")
    }

    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        return nil, fmt.Errorf("DB_NAME environment variable not set")
    }

    dbHost := os.Getenv("DB_HOST")
    if dbHost == "" {
        return nil, fmt.Errorf("DB_HOST environment variable not set")
    }

    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        return nil, fmt.Errorf("failed to open MySQL connection: %w", err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
    }

    fmt.Println("Connected successfully")

    return db, nil
}
