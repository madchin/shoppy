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
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Get(uuid string) (user.User, error) {
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
func (ur *UserRepository) UpdateName(uuid string, name string, updateFn func(user.User) (user.User, error)) error {
	return nil
}
func (ur *UserRepository) UpdateEmail(uuid string, email string, updateFn func(user.User) (user.User, error)) error {
	return nil
}

func (ur *UserRepository) Delete(uuid string) error {
	return nil
}
