package handlers

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "strconv"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/samiyonas/to-do-API/models"
)

func Task_by_Id(w http.ResponseWriter, r *http.Request) {
    task_id := chi.URLParam(r, "userID")

    taskId, err := strconv.Atoi(task_id)
    if err != nil {
        http.Error(w, "Error retrieveing URL Param\n", http.StatusInternalServerError)
    }

    task, err := models.Task_by_Id(taskId)
    if err != nil {
        if err == sql.ErrNoRows{
            http.Error(w, "User not found\n", http.StatusNotFound)
            return
        }
        http.Error(w, fmt.Sprintf("Error querying task %v\n", err), http.StatusInternalServerError)
    }

    jsonData, err := json.Marshal(task)
    if err != nil {
        http.Error(w, "couldn't decode your task\n", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}
