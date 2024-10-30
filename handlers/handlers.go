package handlers

import (
    "fmt"
    "encoding/json"
    "net/http"
    "github.com/samiyonas/to-do-API/models"
)

func Add_user(w http.ResponseWriter, r *http.Request) {
    var user models.User

    if r.Method != http.MethodPost {
        http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
        return
    }

    decoder := json.NewDecoder(r.Body)
    defer r.Body.Close()

    err := decoder.Decode(&user)
    if err != nil {
        http.Error(w, "could not decode user info", http.StatusBadRequest)
        return
    }

    err = models.Add_user(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "You've been added!")
}

func Add_task(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    if r.Method != http.MethodPost {
        http.Error(w, "Unsupported request method", http.StatusMethodNotAllowed)
        return
    }

    decoder := json.NewDecoder(r.Body)
    defer r.Body.Close()

    err := decoder.Decode(&task)
    if err != nil {
        http.Error(w, "could not decode the task", http.StatusBadRequest)
        return
    }

    err = models.Add_task(task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Your task has been added!")
}
