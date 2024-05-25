package adapters

import (
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type userDTO struct {
	uuid  string
	name  string
	email string
}

type UserRepository struct {
	db *sql.DB
}

var ErrUserNotFound = errors.New("user with provided uuid not found")

func NewUserRepository(db *sql.DB) user.Repository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Get(uuid string) (user.User, custom_error.ContextError) {
	userDto := userDTO{}
	err := ur.db.QueryRow("SELECT * FROM Users WHERE uuid=?", uuid).Scan(&userDto.uuid, &userDto.name, &userDto.email)
	if err != nil {
		return user.User{}, custom_error.NewContextError("user retrieve", custom_error.ErrorTypePersistence, []error{err})
	}
	domainUser := user.New(userDto.name, userDto.email)
	return domainUser, custom_error.ContextError{}
}
func (ur *UserRepository) Create(uuid string, u user.User, createFn func(user.User) (user.User, []error)) custom_error.ContextError {
	return custom_error.ContextError{}
}
func (ur *UserRepository) UpdateName(uuid string, name string, updateFn func(user.User) (user.User, []error)) custom_error.ContextError {
	return custom_error.ContextError{}
}
func (ur *UserRepository) UpdateEmail(uuid string, email string, updateFn func(user.User) (user.User, []error)) custom_error.ContextError {
	return custom_error.ContextError{}
}

func (ur *UserRepository) Delete(uuid string) custom_error.ContextError {
	return custom_error.ContextError{}
}
