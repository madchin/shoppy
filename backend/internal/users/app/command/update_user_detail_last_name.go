package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type UpdateUserDetailLastNameHandler decorator.CommandHandler[updateUserDetailLastName]

type updateUserDetailLastName struct {
	uuid     string
	lastName string
}

type updateUserDetailLastNameHandler struct {
	repo user.DetailRepository
}

func NewUpdateUserDetailLastName(uuid string, lastName string) updateUserDetailLastName {
	return updateUserDetailLastName{uuid, lastName}
}

func NewUpdateUserDetailLastNameHandler(repo user.DetailRepository, logger *logrus.Entry) UpdateUserDetailLastNameHandler {
	return decorator.ApplyCommandHandler(updateUserDetailLastNameHandler{repo}, logger)
}

func (u updateUserDetailLastNameHandler) Handle(cmd updateUserDetailLastName) custom_error.ContextError {
	return u.repo.UpdateLastName(cmd.uuid, cmd.lastName, func(u user.UserDetail) error {
		return u.ValidateLastName()
	})
}
