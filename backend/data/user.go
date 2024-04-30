package data

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repo *UserRepository) GetUser(uuid string) (*User, error) {
	if uuid == "" {
		return nil, &ErrMissingUuid{}
	}
	user := &User{}
	err := repo.Db.QueryRow("SELECT * FROM Users WHERE uuid=?", uuid).Scan(&user.Uuid, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Create(user *User) error {
	if user.Uuid == "" {
		user.Uuid = uuid.New().String()
	}
	if user.Email == "" {
		return &ErrEmptyEmail{user: user}
	}
	_, err := repo.Db.Exec("INSERT INTO Users (uuid, name, email) VALUES (?, ?, ?)", user.Uuid, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) UpdateName(user *User) error {
	if user.Uuid == "" {
		return &ErrMissingUuid{user: user}
	}
	_, err := repo.Db.Exec("UPDATE Users SET name=? WHERE uuid=?", user.Name, user.Uuid)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) UpdateEmail(user *User) error {
	if user.Uuid == "" {
		return &ErrMissingUuid{user: user}
	}
	if user.Email == "" {
		return &ErrEmptyEmail{user: user}
	}
	_, err := repo.Db.Exec("UPDATE Users SET email=? WHERE uuid=?", user.Email, user.Uuid)
	if err != nil {
		return err
	}
	return nil
}
