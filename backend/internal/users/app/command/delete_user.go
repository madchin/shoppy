package command

import (
	"backend/internal/common/decorator"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type DeleteUserHandler decorator.CommandHandler[deleteUser]

type deleteUser struct {
	uuid string
}

type deleteUserHandler struct {
	repo user.Repository
}

func NewDeleteUser(uuid string) deleteUser {
	return deleteUser{uuid}
}

func NewDeleteUserHandler(repo user.Repository, logger *logrus.Entry) DeleteUserHandler {
	return decorator.ApplyCommandHandler(deleteUserHandler{repo}, logger)
}

func (u deleteUserHandler) Handle(cmd deleteUser) error {
	return u.repo.Delete(cmd.uuid)
}
