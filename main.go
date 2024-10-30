package main

import (
    "fmt"
    "net/http"
    "github.com/samiyonas/to-do-API/handlers"
    "github.com/samiyonas/to-do-API/models"
)



func main() {
    Db, err := models.Init_db()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    err = Db.Ping()
    if err != nil {
        fmt.Println("Error connecting to the database:", err)
    }

    defer Db.Close()

    err = models.Create_tables()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    http.HandleFunc("/register", handlers.Add_user)
    http.HandleFunc("/addtask", handlers.Add_task)

    http.ListenAndServe(":8080", nil)
}
