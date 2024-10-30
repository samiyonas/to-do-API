package models

import (
    "errors"
)

func Add_user(user User) (error) {
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
    _, err = db.Exec(cmd, user.Name, user.Email)
    if err != nil {
        return errors.New("not inserted")
    }

    return nil
}
