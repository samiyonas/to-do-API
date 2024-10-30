package handlers

func Task_by_Id(w http.ResponseWriter, r *http.Request) {
    task_id := chi.URLParam(r, "userID")

    task, err := models.Task_by_Id(task_id)
    if err != nil {
        if err == sql.ErrNoRows{
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }
        http.Error(w, fmt.Sprintf("Error querying task %v", err), http.StatusInternalServerError)
    }

    jsonData, err := json.Marshal(task)
    if err != nil {
        http.Error(w, "couldn't decode your task", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}
