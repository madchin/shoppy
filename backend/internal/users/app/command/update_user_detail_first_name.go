package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type UpdateUserDetailFirstNameHandler decorator.CommandHandler[updateUserDetailFirstName]

type updateUserDetailFirstName struct {
	uuid      string
	firstName string
}

type updateUserDetailFirstNameHandler struct {
	repo user.DetailRepository
}

func NewUpdateUserDetailFirstName(uuid string, firstName string) updateUserDetailFirstName {
	return updateUserDetailFirstName{uuid, firstName}
}

func NewUpdateUserDetailFirstNameHandler(repo user.DetailRepository, logger *logrus.Entry) UpdateUserDetailFirstNameHandler {
	return decorator.ApplyCommandHandler(updateUserDetailFirstNameHandler{repo}, logger)
}

func (u updateUserDetailFirstNameHandler) Handle(cmd updateUserDetailFirstName) custom_error.ContextError {
	return u.repo.UpdateFirstName(cmd.uuid, cmd.firstName, func(u user.UserDetail) (user.UserDetail, error) {
		err := u.ValidateFirstName()
		if err != nil {
			return user.UserDetail{}, err
		}
		return u, nil
	})
}
