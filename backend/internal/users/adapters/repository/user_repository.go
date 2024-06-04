package adapters

import (
	common_adapter "backend/internal/common/adapters"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type userDTO struct {
	uuid     string
	name     string
	email    string
	password string
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
	err := ur.db.QueryRow("SELECT * FROM Users WHERE uuid=?", uuid).Scan(&userDto.uuid, &userDto.name, &userDto.email, &userDto.password)
	if err == sql.ErrNoRows {
		return user.User{}, custom_error.NewPersistenceError("user retrieve", "user with provided uuid not found")
	}

	if err != nil {
		return user.User{}, custom_error.UnknownPersistenceError("user retrieve")
	}

	domainUser := user.NewUser(userDto.name, userDto.email, userDto.password)
	return domainUser, custom_error.ContextError{}
}
func (ur *UserRepository) Create(uuid string, u user.User, createFn func(user.User) []error) custom_error.ContextError {
	errs := createFn(u)
	if len(errs) > 0 {
		return custom_error.NewValidationErrors("user add", errs)
	}

	if _, err := ur.db.Exec("INSERT INTO Users (uuid, name, email, password) VALUES (?, ?, ?, ?)", uuid, u.Name(), u.Email(), u.Password()); err != nil {
		if common_adapter.IsDuplicateEntryError(err) {
			return custom_error.NewPersistenceError("user add", "user with provided email already exists")
		}
		return custom_error.UnknownPersistenceError("user add")
	}

	return custom_error.ContextError{}
}
func (ur *UserRepository) UpdateName(uuid string, name string, updateFn func(user.User) []error) custom_error.ContextError {
	u, err := ur.Get(uuid)
	if err.Error() != "" {
		return err
	}

	errs := updateFn(u)
	if len(errs) > 0 {
		return custom_error.NewValidationErrors("user update name", errs)
	}

	if _, err := ur.db.Exec("UPDATE Users SET name=? WHERE uuid=?", u.Name(), uuid); err != nil {
		return custom_error.UnknownPersistenceError("user update name")
	}

	return custom_error.ContextError{}
}
func (ur *UserRepository) UpdateEmail(uuid string, email string, updateFn func(user.User) []error) custom_error.ContextError {
	u, err := ur.Get(uuid)
	if err.Error() != "" {
		return err
	}
	errs := updateFn(u)
	if len(errs) > 0 {
		return custom_error.NewValidationErrors("user update email", errs)
	}

	if _, err := ur.db.Exec("UPDATE Users SET email=? WHERE uuid=?", u.Email(), uuid); err != nil {
		return custom_error.UnknownPersistenceError("user update email")
	}

	return custom_error.ContextError{}
}

func (ur *UserRepository) UpdatePassword(uuid string, password string, updateFn func(user.User) []error) custom_error.ContextError {
	u, err := ur.Get(uuid)
	if err.Error() != "" {
		return err
	}

	errs := updateFn(u)
	if len(errs) > 0 {
		return custom_error.NewValidationErrors("user update password", errs)
	}

	if _, err := ur.db.Exec("UPDATE Users SET password=? WHERE uuid=?", u.Password(), uuid); err != nil {
		return custom_error.UnknownPersistenceError("user update password")
	}

	return custom_error.ContextError{}
}

func (ur *UserRepository) Delete(uuid string) custom_error.ContextError {
	if _, err := ur.Get(uuid); err.Error() != "" {
		return err
	}

	_, err := ur.db.Exec("DELETE FROM Users WHERE uuid=?", uuid)
	if err != nil {
		return custom_error.UnknownPersistenceError("user delete")
	}

	return custom_error.ContextError{}
}
