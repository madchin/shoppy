package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"context"

	"github.com/sirupsen/logrus"
)

type DeleteUserDetailHandler decorator.CommandHandler[deleteUserDetail]

type deleteUserDetail struct {
	userUuid string
}

type deleteUserDetailHandler struct {
	repo user.DetailRepository
}

func NewDeleteUserDetail(userUuid string) deleteUserDetail {
	return deleteUserDetail{userUuid}
}

func NewDeleteUserDetailHandler(repo user.DetailRepository, logger *logrus.Entry) DeleteUserDetailHandler {
	return decorator.ApplyCommandHandler(deleteUserDetailHandler{repo}, logger)
}

func (u deleteUserDetailHandler) Handle(ctx context.Context, cmd deleteUserDetail) custom_error.ContextError {
	return u.repo.Delete(ctx, cmd.userUuid)
}
