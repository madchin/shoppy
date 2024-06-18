package query

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"context"

	"github.com/sirupsen/logrus"
)

type loginUser struct {
	email    string
	password string
}

type loginUserHandler struct {
	repo user.Repository
}

type LoginUserHandler decorator.QueryHandler[loginUser, user.User]

func NewLoginUser(email string, password string) loginUser {
	return loginUser{email, password}
}

func NewLoginUserHandler(repo user.Repository, logger *logrus.Entry) LoginUserHandler {
	return decorator.ApplyQueryHandler(loginUserHandler{repo}, logger)
}

func (luh loginUserHandler) Handle(ctx context.Context, q loginUser) (user.User, custom_error.ContextError) {
	return luh.repo.GetByEmail(ctx, q.email, func(u user.User) error {
		return u.IsPasswordEqual(q.password)
	})
}
