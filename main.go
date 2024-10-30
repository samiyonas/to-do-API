package main

import (
    "net/http"
    "github.com/samiyonas/to-do-API/handlers"
    "github.com/samiyonas/to-do-API/models"
)



func main() {
    _, err := models.init_db()
    if err != nil {
        fmt.Errorf(err)
        return
    }

    defer db.Close()

    err = create_tables()
    if err != nil {
        fmt.Errorf(err)
        return
    }

    http.HandleFunc("/register", handlers.add_user)
    http.HandleFunc("/addtask", handlers.add_task)

    http.ListenAndServe(":8080", nil)
}
