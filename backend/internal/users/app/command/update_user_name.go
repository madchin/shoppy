package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"context"

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

func (u updateUserNameHandler) Handle(ctx context.Context, cmd updateUserName) custom_error.ContextError {
	return u.repo.UpdateName(ctx, cmd.uuid, cmd.name, func(u user.User) []error {
		return u.ValidateName()
	})
}
