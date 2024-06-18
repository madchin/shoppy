package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"context"

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

func (u updateUserEmailHandler) Handle(ctx context.Context, cmd updateUserEmail) custom_error.ContextError {
	return u.repo.UpdateEmail(ctx, cmd.uuid, cmd.email, func(u user.User) []error {
		return u.ValidateEmail()
	})
}
