package handlers

import (
    "encoding/json"
    "net/http"
    "models"
)

func add_user(w http.ResponseWriter, r *http.Request) {
    var user User

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

    err = models.add_user(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "You've been added!")
}

func add_task(w http.ResponseWriter, r *http.Request) {
    var task Task
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

    err = models.add_task(task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Your task has been added!")
}
