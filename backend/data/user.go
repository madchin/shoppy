package data

import (
	"database/sql"

	"github.com/google/uuid"
)

func GetUser(db *sql.DB, uuid string) (*User, error) {
	user := &User{}
	err := db.QueryRow("SELECT * FROM Users WHERE uuid=?", uuid).Scan(&user.Uuid, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) Create(db *sql.DB) error {
	if user.Uuid == "" {
		user.Uuid = uuid.New().String()
	}
	if user.Email == "" {
		return Err.EmptyEmail
	}
	_, err := db.Exec("INSERT INTO Users (uuid, name, email) VALUES (?, ?, ?)", user.Uuid, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) UpdateName(db *sql.DB) error {
	_, err := db.Exec("UPDATE Users SET name=? WHERE uuid=?", user.Name, user.Uuid)
	if err != nil {
		return err
	}
	return nil
}

func (user *User) UpdateEmail(db *sql.DB) error {
	_, err := db.Exec("UPDATE Users SET email=? WHERE uuid=?", user.Email, user.Uuid)
	if err != nil {
		return err
	}
	return nil
}
