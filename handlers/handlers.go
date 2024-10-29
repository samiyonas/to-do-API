package handlers

import (
    "encoding/json"
    "net/http"
    "models"
)

func add_task(w http.ResponseWriter, r *http.Request) {
    var task Task
    if r.Method != http.MethodPost {
        fmt.Fprintf(w, "Unsupported request method", http.StatusMethodNotAllowed)
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
        http.Error(w, "could not add the task", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "Your task has been added!")
}
