package data

import (
	"database/sql"

	"github.com/google/uuid"
)

type userRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{Db: db}
}

func (repo *userRepository) GetUser(uuid string) (User, error) {
	user := User{}
	if uuid == "" {
		return User{}, ErrMissingUuid{}
	}
	err := repo.Db.QueryRow("SELECT * FROM Users WHERE uuid=?", uuid).Scan(&user.Uuid, &user.Name, &user.Email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (repo *userRepository) Create(user User) error {
	if user.Uuid == "" {
		user.Uuid = uuid.New().String()
	}
	if user.Email == "" {
		return ErrEmptyEmail{user: user}
	}
	_, err := repo.Db.Exec("INSERT INTO Users (uuid, name, email) VALUES (?, ?, ?)", user.Uuid, user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) UpdateName(user User) error {
	if user.Uuid == "" {
		return ErrMissingUuid{}
	}
	_, err := repo.Db.Exec("UPDATE Users SET name=? WHERE uuid=?", user.Name, user.Uuid)
	if err != nil {
		return err
	}
	return nil
}

func (repo *userRepository) UpdateEmail(user User) error {
	if user.Uuid == "" {
		return ErrMissingUuid{}
	}
	if user.Email == "" {
		return ErrEmptyEmail{user: user}
	}
	_, err := repo.Db.Exec("UPDATE Users SET email=? WHERE uuid=?", user.Email, user.Uuid)
	if err != nil {
		return err
	}
	return nil
}
