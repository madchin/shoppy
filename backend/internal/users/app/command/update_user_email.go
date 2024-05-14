package command

import (
	"backend/internal/common/decorator"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type UpdateUserEmailHandler decorator.CommandHandler[updateUserEmail]

type updateUserEmail struct {
	uuid  string
	email string
}

type updateUserEmailHandler struct {
	repo user.Repository
}

func NewUpdateUserEmail(uuid string, email string) updateUserEmail {
	return updateUserEmail{uuid, email}
}

func NewUpdateUserEmailHandler(repo user.Repository, logger *logrus.Entry) UpdateUserEmailHandler {
	return decorator.ApplyCommandHandler(updateUserEmailHandler{repo}, logger)
}

func (u updateUserEmailHandler) Handle(cmd updateUserEmail) error {
	return u.repo.UpdateEmail(cmd.uuid, cmd.email, func(u user.User) (user.User, error) {
		err := u.ValidateEmail()
		if err != nil {
			return user.User{}, err
		}
		return u, nil
	})
}
