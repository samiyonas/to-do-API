package models

func All_tasks(user_id int) ([]Task, error) {
    var tasks []Task
    query := "SELECT * FROM todo WHERE user_id = ?"
    
    result, err := Db.Query(query, user_id)
    if err != nil {
        return []Task{}, err
    }

    defer result.Close()
    for result.Next() {
        var task Task 
        var unusedId interface{}

        err := result.Scan(&task.User_id, &unusedId, &task.Title, &task.Content, &task.Done)
        if err != nil {
            return []Task{}, err
        }

        tasks = append(tasks, task)
    }

    return tasks, nil
}
