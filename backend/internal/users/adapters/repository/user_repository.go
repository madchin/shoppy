package adapters

import (
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"database/sql"
	"errors"
	"strings"

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
		return user.User{}, custom_error.NewContextError("user retrieve", custom_error.ErrorTypePersistence, []error{errors.New("user with provided uuid not found")})
	}

	if err != nil {
		return user.User{}, unknownPersistenceError("user retrieve")
	}

	domainUser := user.New(userDto.name, userDto.email, userDto.password)
	return domainUser, custom_error.ContextError{}
}
func (ur *UserRepository) Create(uuid string, u user.User, createFn func(user.User) (user.User, []error)) custom_error.ContextError {
	u, errs := createFn(u)
	if len(errs) > 0 {
		return custom_error.NewContextError("user add", custom_error.ErrorTypeValidation, errs)
	}

	if _, err := ur.db.Exec("INSERT INTO Users (uuid, name, email, password) VALUES (?, ?, ?, ?)", uuid, u.Name(), u.Email(), u.Password()); err != nil {
		if isDuplicateEntryError(err) {
			return custom_error.NewContextError("user add", custom_error.ErrorTypePersistence, []error{errors.New("user with provided email already exists")})
		}
		return unknownPersistenceError("user add")
	}

	return custom_error.ContextError{}
}
func (ur *UserRepository) UpdateName(uuid string, name string, updateFn func(user.User) (user.User, []error)) custom_error.ContextError {
	u, err := ur.Get(uuid)
	if err.Error() != "" {
		return err
	}

	u, errs := updateFn(u)
	if len(errs) > 0 {
		return custom_error.NewContextError("user update name", custom_error.ErrorTypeValidation, errs)
	}

	if _, err := ur.db.Exec("UPDATE Users SET name=? WHERE uuid=?", u.Name(), uuid); err != nil {
		return unknownPersistenceError("user update name")
	}

	return custom_error.ContextError{}
}
func (ur *UserRepository) UpdateEmail(uuid string, email string, updateFn func(user.User) (user.User, []error)) custom_error.ContextError {
	u, err := ur.Get(uuid)
	if err.Error() != "" {
		return err
	}

	u, errs := updateFn(u)
	if len(errs) > 0 {
		return custom_error.NewContextError("user update email", custom_error.ErrorTypePersistence, errs)
	}

	if _, err := ur.db.Exec("UPDATE Users SET email=? WHERE uuid=?", u.Email(), uuid); err != nil {
		return unknownPersistenceError("user update email")
	}

	return custom_error.ContextError{}
}

func (ur *UserRepository) UpdatePassword(uuid string, password string, updateFn func(user.User) (user.User, []error)) custom_error.ContextError {
	u, err := ur.Get(uuid)
	if err.Error() != "" {
		return err
	}

	u, errs := updateFn(u)
	if len(errs) > 0 {
		return custom_error.NewContextError("user update password", custom_error.ErrorTypePersistence, errs)
	}

	if _, err := ur.db.Exec("UPDATE Users SET password=? WHERE uuid=?", u.Password(), uuid); err != nil {
		return unknownPersistenceError("user update password")
	}

	return custom_error.ContextError{}
}

func (ur *UserRepository) Delete(uuid string) custom_error.ContextError {
	if _, err := ur.Get(uuid); err.Error() != "" {
		return err
	}

	_, err := ur.db.Exec("DELETE FROM Users WHERE uuid=?", uuid)
	if err != nil {
		return unknownPersistenceError("user delete")
	}

	return custom_error.ContextError{}
}

func isDuplicateEntryError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "duplicate")
}

func unknownPersistenceError(context string) custom_error.ContextError {
	return custom_error.NewContextError(context, custom_error.ErrorTypePersistence, []error{errors.New("unknown")})
}
