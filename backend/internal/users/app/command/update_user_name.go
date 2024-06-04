package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type UpdateUserNameHandler decorator.CommandHandler[updateUserName]

type updateUserName struct {
	uuid string
	name string
}

type updateUserNameHandler struct {
	repo user.Repository
}

func NewUpdateUserName(uuid string, name string) updateUserName {
	return updateUserName{uuid, name}
}

func NewUpdateUserNameHandler(repo user.Repository, logger *logrus.Entry) UpdateUserNameHandler {
	return decorator.ApplyCommandHandler(updateUserNameHandler{repo}, logger)
}

func (u updateUserNameHandler) Handle(cmd updateUserName) custom_error.ContextError {
	return u.repo.UpdateName(cmd.uuid, cmd.name, func(u user.User) []error {
		return u.ValidateName()
	})
}
