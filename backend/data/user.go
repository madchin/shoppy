package data

import (
	"database/sql"

	"github.com/google/uuid"
)

type User struct {
	Email string `json:"email"`
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
}

func GetUser(db *sql.DB, uuid string) (*User, error) {
	user := &User{}
	err := db.QueryRow("SELECT * FROM Users WHERE uuid=$1", uuid).Scan(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(db *sql.DB, user *User) error {
	uuid := uuid.New().String()
	_, err := db.Exec("INSERT INTO Users (uuid, name, email) VALUES (?, ?, ?)", uuid, user.Name, user.Email)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT * FROM Users WHERE uuid=?", uuid).Scan(&user.Uuid, &user.Email, &user.Name)
	if err != nil {
		return err
	}
	return nil
}
