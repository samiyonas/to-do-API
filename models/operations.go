package models

import (
    "errors"
)

func Add_user(user User) (int64, error) {
    query := "SELECT * FROM user WHERE email = ?"
    result, err := Db.Query(query, user.Email)
    if err != nil {
        return -1, errors.New("this is internal failure")
    }
    if result.Next() {
        result.Close()
        return -1, errors.New("email already exists")
    }
    cmd := "INSERT INTO user (name, email) VALUES (?, ?)"
    info, err := Db.Exec(cmd, user.Name, user.Email)
    if err != nil {
        return -1, errors.New(err.Error())
    }
    id, err := info.LastInsertId()
    if err != nil {
        return -1, errors.New(err.Error())
    }

    return id, nil
}

func Add_task(task Task) (int64, error) {
    query := "SELECT * FROM user WHERE id = ?"
    result, err := Db.Query(query, task.User_id)
    if err != nil {
        return -1, errors.New(err.Error())
    }
    if !result.Next() {
        return -1, errors.New(err.Error())
    }

    result.Close()

    cmd := "INSERT INTO todo (user_id, title, content, done) VALUES(?, ?, ?, ?)"

    info, err := Db.Exec(cmd, task.User_id, task.Title, task.Content, task.Done)
    if err != nil {
        return -1, errors.New("task not added")
    }
    id, err := info.LastInsertId()
    if err != nil {
        return -1, errors.New("not the right ID")
    }

    return id, nil
}

func Task_by_Id(id int) (Task, error) {
    var task Task

    query := "SELECT * FROM todo WHERE id = ?"

    err := Db.QueryRow(query, id).Scan(&task.User_id, &task.Title, &task.Content, &task.Done)
    if err != nil {
        return task, err
    }

    return task, nil
}
