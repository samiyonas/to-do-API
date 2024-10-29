package models

func add_user(user User) (error) {
    query := "INSERT INTO (name, email) VALUES (?, ?)"
    _, err := db.Exec(query, user.name, user.email)
    if err != nil {
        return errors.New("not inserted")
    }

    return nil
}
