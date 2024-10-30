package main

// Entry point of the API

import (
    "fmt" // For formatting string
    "net/http" // For http connections and routing
    "github.com/samiyonas/to-do-API/handlers" // For handling routers
    "github.com/samiyonas/to-do-API/models" // For Database related functions
    "github.com/go-chi/chi/v5" // Package for complex routing
)

func main() {
    // database initialization and error handling
    Db, err := models.Init_db()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // checking database connection and if not connected error is handled
    err = Db.Ping()
    if err != nil {
        fmt.Println("Error connecting to the database:", err)
    }

    // closing the DB connection gracefully if the server crashes
    defer Db.Close()

    // create the tables if they don't exist and error handling
    err = models.Create_tables()
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // chi router object
    r := chi.NewRouter()

    // register different routes to our router
    r.Post("/register", handlers.Add_user)
    r.Post("/addtask", handlers.Add_task)
    r.Get("/taskbyid/{taskID}", handlers.Task_by_Id)
    r.Get("/alltasks/{userID}", handlers.All_tasks)

    // start serving on port 8080 localhost using chi object
    http.ListenAndServe(":8080", r)
}
