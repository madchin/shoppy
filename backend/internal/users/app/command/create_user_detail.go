package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type createUserDetail struct {
	uuid string
	user user.UserDetail
}

type CreateUserDetailHandler decorator.CommandHandler[createUserDetail]

type createUserDetailHandler struct {
	repo user.DetailRepository
}

func NewCreateUserDetail(uuid string, user user.UserDetail) createUserDetail {
	return createUserDetail{uuid, user}
}

func NewCreateUserDetailHandler(repo user.DetailRepository, logger *logrus.Entry) decorator.CommandHandler[createUserDetail] {
	return decorator.ApplyCommandHandler(createUserDetailHandler{repo}, logger)
}

func (ru createUserDetailHandler) Handle(cmd createUserDetail) custom_error.ContextError {
	return ru.repo.Create(cmd.uuid, cmd.user, func(u user.UserDetail) (user.UserDetail, []error) {
		errs := u.Validate()
		if errs != nil {
			return user.UserDetail{}, errs
		}
		return u, nil
	})

}
