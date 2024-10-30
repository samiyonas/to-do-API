package handlers

import (
    "fmt" // For formatting string
    "encoding/json" // For encoding and decoding json
    "net/http" // For handling requests
    "github.com/samiyonas/to-do-API/models" // To retrieve from database
)

func Add_user(w http.ResponseWriter, r *http.Request) {
    // user object
    var user models.User

    // check for method similarity
    if r.Method != http.MethodPost {
        http.Error(w, "Unsupported request method\n", http.StatusMethodNotAllowed)
        return
    }

    // decode the request body into our struct type
    decoder := json.NewDecoder(r.Body)
    defer r.Body.Close()

    err := decoder.Decode(&user)
    if err != nil {
        http.Error(w, "could not decode user info\n", http.StatusBadRequest)
        return
    }

    // Add the user to the database
    id, err := models.Add_user(user)
    if err != nil {
        http.Error(w, fmt.Sprintf("couldn't add you\n", err), http.StatusInternalServerError)
        return
    }

    // preparing the response body
    ret := fmt.Sprintf("You've been added and your ID is %d\n", id)

    // Writes it to the response body
    w.Write([]byte(ret))
}

func Add_task(w http.ResponseWriter, r *http.Request) {
    // task table object
    var task models.Task
    // checking for method similarity
    if r.Method != http.MethodPost {
        http.Error(w, "Unsupported request method\n", http.StatusMethodNotAllowed)
        return
    }

    // decoding request body to our type struct
    decoder := json.NewDecoder(r.Body)
    defer r.Body.Close()

    err := decoder.Decode(&task)
    if err != nil {
        http.Error(w, "could not decode the task\n", http.StatusBadRequest)
        return
    }

    // Add task to the database
    id, err := models.Add_task(task)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    // prepare the response body
    ret := fmt.Sprintf("Your task has been added! You're task ID is %d\n", id)
    // write to the response body
    w.Write([]byte(ret))
}
