package data

import (
	"database/sql"

	"github.com/google/uuid"
)

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

func UpdateUserName(db *sql.DB, user *User) error {
	_, err := db.Exec("UPDATE Users SET name=?;", user.Name)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserEmail(db *sql.DB, user *User) error {
	_, err := db.Exec("UPDATE Users SET email=?;", user.Email)
	if err != nil {
		return err
	}
	return nil
}
