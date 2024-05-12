package adapters

import (
	"backend/internal/common/repository"
	"backend/internal/users/domain/user"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type userDTO struct {
	uuid  string
	name  string
	email string
}

type UserRepository struct {
	db   *sql.DB
	repo user.Repository
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Get(uuid string, getFn func(user.User) (user.User, error)) (user.User, error) {
	userDto := userDTO{}
	err := ur.db.QueryRow("SELECT * FROM Users WHERE uuid=?", uuid).Scan(&userDto.uuid, &userDto.name, &userDto.email)
	if err != nil {
		return user.User{}, repository.ErrInternal
	}
	domainUser := user.New(userDto.name, userDto.email)
	return domainUser, nil
}
func (ur *UserRepository) Create(uuid string, u user.User, createFn func(user.User) (user.User, error)) error {
	return nil
}
func (ur *UserRepository) UpdateName(uuid string, u user.User, updateFn func(user.User) (user.User, error)) error {
	return nil
}
func (ur *UserRepository) UpdateEmail(uuid string, u user.User, updateFn func(user.User) (user.User, error)) error {
	return nil
}

func (ur *UserRepository) Delete(uuid string, deleteFn func(user.User) error) error {
	return nil
}

// Create(
// 	uuid string,
// 	user User,
// 	createFn func(User) (User, error),
// ) error
// Get(
// 	uuid string,
// 	getFn func(User) (User, error),
// ) (User, error)
// UpdateName(
// 	uuid string,
// 	user User,
// 	updateFn func(User) (User, error),
// ) error
// UpdateEmail(
// 	uuid string,
// 	user User,
// 	updateFn func(User) (User, error),
// ) error
// Delete(
// 	uuid string,
// 	user User,
// 	deleteFn func(User) error,
// ) error
