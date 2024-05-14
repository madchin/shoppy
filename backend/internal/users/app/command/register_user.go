package command

import (
	"backend/internal/common/decorator"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type registerUser struct {
	uuid string
	user user.User
}

type RegisterUserHandler decorator.CommandHandler[registerUser]

type registerUserHandler struct {
	repo user.Repository
}

func NewRegisterUser(uuid string, user user.User) registerUser {
	return registerUser{uuid, user}
}

func NewRegisterUserHandler(repo user.Repository, logger *logrus.Entry) decorator.CommandHandler[registerUser] {
	return decorator.ApplyCommandHandler(registerUserHandler{repo}, logger)
}

func (ru registerUserHandler) Handle(cmd registerUser) error {
	return ru.repo.Create(cmd.uuid, cmd.user, func(u user.User) (user.User, error) {
		err := u.Validate()
		if err != nil {
			return user.User{}, err
		}
		return u, nil
	})
}
