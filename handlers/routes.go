package handlers

import (
    "database/sql" // For database connections
    "encoding/json" // For encoding and decoding json
    "fmt" // For formatting string
    "strconv" // For changing string to other data types
    "net/http" // For handling web request and response
    "github.com/go-chi/chi/v5" // Package for complex routing
    "github.com/samiyonas/to-do-API/models" // For database operations
)

func Task_by_Id(w http.ResponseWriter, r *http.Request) {
    // Get the URL Param
    task_id := chi.URLParam(r, "taskID")

    // Change the URL Param from string to int
    taskId, err := strconv.Atoi(task_id)
    if err != nil {
        http.Error(w, "Error retrieveing URL Param\n", http.StatusInternalServerError)
    }

    // Get the task by ID
    task, err := models.Task_by_Id(taskId)
    if err != nil {
        if err == sql.ErrNoRows{
            http.Error(w, "User not found\n", http.StatusNotFound)
            return
        }
        http.Error(w, fmt.Sprintf("Error querying task %v\n", err), http.StatusInternalServerError)
    }

    // encode our struct type to a JSON
    jsonData, err := json.Marshal(task)
    if err != nil {
        http.Error(w, "couldn't decode your task\n", http.StatusInternalServerError)
        return
    }

    // Set the content type to JSON
    w.Header().Set("Content-Type", "application/json")
    // Write the JSON to our response type
    w.Write(jsonData)
}

func All_tasks(w http.ResponseWriter, r *http.Request) {
    // Get the URL Param
    user_id := chi.URLParam(r, "userID")

    // Change the URL Param from string to int
    userId, err := strconv.Atoi(user_id)
    if err != nil {
        http.Error(w, "Error retrieving URL Param\n", http.StatusInternalServerError)
        return
    }
    
    // Get all tasks from database by ID
    tasks, err := models.All_tasks(userId)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Change the tasks to JSON
    jsonData, err := json.Marshal(tasks)
    if err != nil {
        http.Error(w, "couldn't decode your task\n", http.StatusInternalServerError)
        return
    }

    // Set the content type to JSON
    w.Header().Set("Content-Type", "application/json")
    // Write the JSON to our response type
    w.Write(jsonData)
}
