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
        http.Error(w, "Unsupported request method\n", http.StatusMethodNotAllowed)
        return
    }

    decoder := json.NewDecoder(r.Body)
    defer r.Body.Close()

    err := decoder.Decode(&user)
    if err != nil {
        http.Error(w, "could not decode user info\n", http.StatusBadRequest)
        return
    }

    id, err := models.Add_user(user)
    if err != nil {
        http.Error(w, fmt.Sprintf("couldn't add you\n", err), http.StatusInternalServerError)
        return
    }

    ret := fmt.Sprintf("You've been added and your ID is %d\n", id)

    w.Write([]byte(ret))
}

func Add_task(w http.ResponseWriter, r *http.Request) {
    var task models.Task
    if r.Method != http.MethodPost {
        http.Error(w, "Unsupported request method\n", http.StatusMethodNotAllowed)
        return
    }

    decoder := json.NewDecoder(r.Body)
    defer r.Body.Close()

    err := decoder.Decode(&task)
    if err != nil {
        http.Error(w, "could not decode the task\n", http.StatusBadRequest)
        return
    }

    id, err := models.Add_task(task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    ret := fmt.Sprintf("Your task has been added! You're task ID is %d\n", id)
    w.Write([]byte(ret))
}
