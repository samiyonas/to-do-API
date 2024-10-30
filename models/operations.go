package models

import (
    "errors"
)

func Add_user(user User) (int64, error) {
    query := "SELECT * FROM user WHERE email = ?"
    result, err := db.Query(query, user.Email)
    if err != nil {
        return errors.New("this is internal failure")
    }
    if result.Next() {
        result.Close()
        return errors.New("email already exists")
    }
    cmd := "INSERT INTO (name, email) VALUES (?, ?)"
    info, err = db.Exec(cmd, user.Name, user.Email)
    if err != nil {
        return errors.New("not inserted")
    }
    id, err := info.LastInsertedId()
    if err != nil {
        return -1, errors.New("not the right ID")
    }

    return id, nil
}

func Add_task(task Task) (int64, error) {
    cmd := "INSERT INTO todo (user_id, title, content, done) VALUES(?, ?, ?, ?)"

    info, err := db.Exec(cmd, task.User_id, task.Title, task.Content, task.Done)
    if err != nil {
        return errors.New("task not added")
    }
    id, err := info.LastInsertedId()
    if err != nil {
        return -1, errors.New("not the right ID")
    }

    return nil, id
}
