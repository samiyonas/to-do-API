package models

func add_user(user User) (error) {
    query := "SELECT * FROM user WHERE email = ?"
    result, err = db.Query(query, user.email)
    if err != nil {
        return errors.New("this is internal failure")
    }
    if result.Next() {
        return errors.New("email already exists")
    }
    cmd := "INSERT INTO (name, email) VALUES (?, ?)"
    _, err := db.Exec(cmd, user.name, user.email)
    if err != nil {
        return errors.New("not inserted")
    }

    return nil
}
