package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type updateUserPassword struct {
	uuid     string
	password string
}

type UpdateUserPasswordHandler decorator.CommandHandler[updateUserPassword]

type updateUserPasswordHandler struct {
	userRepository user.Repository
}

func NewUpdateUserPassword(uuid string, password string) updateUserPassword {
	return updateUserPassword{uuid, password}
}

func NewUpdateUserPasswordHandler(userRepository user.Repository, logger *logrus.Entry) UpdateUserPasswordHandler {
	return decorator.ApplyCommandHandler(updateUserPasswordHandler{userRepository}, logger)
}

func (c updateUserPasswordHandler) Handle(cmd updateUserPassword) custom_error.ContextError {
	return c.userRepository.UpdatePassword(cmd.uuid, cmd.password, func(u user.User) []error {
		return u.ValidatePassword()
	})
}
