package models

import (
    "fmt" // For formatting strings
    "errors" // For formatting errors
    "database/sql" // For database connections
    "os" // To retrieve data from environment variables
    _ "github.com/go-sql-driver/mysql" // MySQL driver
)

// User information
type User struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

// Task information
type Task struct {
    User_id int `json:"user_id"`
    Title string `json:"title"`
    Content string `json:"content"`
    Done bool `json:"done"`
}

// Database pointer object
var Db *sql.DB

func Init_db() (*sql.DB, error) {
    // database setup
    var err error
    username := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWD")
    database := os.Getenv("DB")
    dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", username, password, database)
    Db, err = sql.Open("mysql", dsn)
    if err != nil {
        return Db, errors.New(err.Error())
    }

    // return the database connection object
    return Db, nil
}

func Create_tables() (error) {
    // create a user table
    user_table := `CREATE TABLE IF NOT EXISTS user (
        id INT AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(50) NOT NULL,
        email VARCHAR(100) NOT NULL UNIQUE
    )`

    // Create the task table
    task_table := `CREATE TABLE IF NOT EXISTS todo (
        user_id INT,
        FOREIGN KEY(user_id) REFERENCES user(id),
        id INT AUTO_INCREMENT UNIQUE,
        title VARCHAR(256),
        content VARCHAR(1024),
        done TINYINT(1) NOT NULL DEFAULT 0
    )`

    // Execute the user table creation
    _, err := Db.Exec(user_table)
    if err != nil {
        return errors.New(err.Error())
    }

    // Execute the Task table creation
    _, err = Db.Exec(task_table)
    if err != nil {
        return errors.New(err.Error())
    }

    return nil
}
