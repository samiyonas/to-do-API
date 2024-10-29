package models

import (
    "fmt"
    "log"
    "database/sql"
    "os"
    _ "github.com/go-sql-driver/mysql"
)

func init_db() (*sql.DB) {
    username := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWD")
    database := os.Getenv("DB")
    dsn := fmt.Sprintf("%s:%s@tcp(localhost:8080)/%s", username, password, database)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error creating the database pointer", err)
    }

    return db
}
