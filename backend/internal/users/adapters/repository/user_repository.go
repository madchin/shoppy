package adapters

import (
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type userDto struct {
	uuid     string
	name     string
	email    string
	password string
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return UserRepository{db: db}
}

func (ur UserRepository) GetByUuid(ctx context.Context, uuid string) (user.User, custom_error.ContextError) {
	userDto := userDto{}
	err := ur.db.QueryRow("SELECT * FROM Users WHERE uuid=?", uuid).Scan(&userDto.uuid, &userDto.name, &userDto.email, &userDto.password)
	if err == sql.ErrNoRows {
		return user.User{}, custom_error.NewPersistenceError("user retrieve", "user with provided uuid not found")
	}

	if err != nil {
		return user.User{}, custom_error.UnknownPersistenceError("user retrieve")
	}

	domainUser := userDto.mapToDomainUser()
	return domainUser, custom_error.ContextError{}
}

func (ur UserRepository) GetByEmail(ctx context.Context, email string, validatePasswordFn func(user.User) error) (user.User, custom_error.ContextError) {
	userDto := userDto{}
	err := ur.db.QueryRow("SELECT * FROM Users WHERE email=?", email).Scan(&userDto.uuid, &userDto.name, &userDto.email, &userDto.password)
	if err == sql.ErrNoRows {
		return user.User{}, custom_error.NewPersistenceError("user retrieve", "user with provided email not found")
	}

	if err != nil {
		return user.User{}, custom_error.UnknownPersistenceError("user retrieve")
	}

	domainUser := userDto.mapToDomainUser()
	err = validatePasswordFn(domainUser)
	if err != nil {
		return user.User{}, custom_error.NewValidationError("user retrieve", err.Error())
	}
	return domainUser, custom_error.ContextError{}
}

func (ur UserRepository) FindByEmail(ctx context.Context, email string) (user.User, custom_error.ContextError) {
	userDto := userDto{}
	err := ur.db.QueryRow("SELECT * FROM Users WHERE email=?", email).Scan(&userDto.uuid, &userDto.name, &userDto.email, &userDto.password)
	if err == sql.ErrNoRows {
		return user.User{}, custom_error.NewPersistenceError("user find", "user with provided email not found")
	}

	if err != nil {
		return user.User{}, custom_error.UnknownPersistenceError("user find")
	}

	domainUser := userDto.mapToDomainUser()
	return domainUser, custom_error.ContextError{}
}

func (ur UserRepository) Create(ctx context.Context, uuid string, u user.User, createFn func(user.User) []error) custom_error.ContextError {
	if user, _ := ur.FindByEmail(ctx, u.Email()); user.Exists() {
		return custom_error.NewPersistenceError("user add", "user with provided email already exists")
	}
	errs := createFn(u)
	if len(errs) > 0 {
		return custom_error.NewValidationErrors("user add", errs)
	}

	if _, err := ur.db.Exec("INSERT INTO Users (uuid, name, email, password) VALUES (?, ?, ?, ?)", uuid, u.Name(), u.Email(), u.Password()); err != nil {
		return custom_error.UnknownPersistenceError("user add")
	}

	return custom_error.ContextError{}
}
func (ur UserRepository) UpdateName(ctx context.Context, uuid string, name string, updateFn func(user.User) []error) custom_error.ContextError {
	u, err := ur.GetByUuid(ctx, uuid)
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
func (ur UserRepository) UpdateEmail(ctx context.Context, uuid string, email string, updateFn func(user.User) []error) custom_error.ContextError {
	u, err := ur.GetByUuid(ctx, uuid)
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

func (ur UserRepository) UpdatePassword(ctx context.Context, uuid string, password string, updateFn func(user.User) []error) custom_error.ContextError {
	u, err := ur.GetByUuid(ctx, uuid)
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

func (ur UserRepository) Delete(ctx context.Context, uuid string) custom_error.ContextError {
	if _, err := ur.GetByUuid(ctx, uuid); err.Error() != "" {
		return err
	}

	_, err := ur.db.Exec("DELETE FROM Users WHERE uuid=?", uuid)
	if err != nil {
		return custom_error.UnknownPersistenceError("user delete")
	}

	return custom_error.ContextError{}
}

func (ur userDto) mapToDomainUser() user.User {
	return user.NewUser(ur.uuid, ur.password, ur.name, ur.email)
}
