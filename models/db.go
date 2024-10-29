package models

import (
    "fmt"
    "errors"
    "log"
    "database/sql"
    "os"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init_db() (*sql.DB, error) {
    username := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWD")
    database := os.Getenv("DB")
    dsn := fmt.Sprintf("%s:%s@tcp(localhost:8080)/%s", username, password, database)
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        return db, errors.New("Error creating the database pointer")
    }

    return db, nil
}

func create_tables() (error) {
    user_table := `CREATE TABLE IF NOT EXISTS user (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(50) NOT NULL,
        email VARCHAR(100) NOT NULL, UNIQUE
    )`

    task_table := `CREATE TABLE IF NOT EXISTS todo (
        user_id INT,
        FOREIGN KEY(user_id) REFERENCES user(id),
        id INT AUTO_INCREMENT UNIQUE,
        title VARCHAR(256),
        content VARCHAR(1024),
        done TINYINT(1) NOT NULL DEFAULT 0
    )`

    _, err := db.Exec(user_table)
    if err != nil {
        return errors.New("error creating table")
    }

    _, err = db.Exec(task_table)
    if err != nil {
        return errors.New("error creating table")
    }

    return nil
}
