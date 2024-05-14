package query

import (
	"backend/internal/common/decorator"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type retrieveUser struct {
	uuid string
}

type RetrieveUserHandler decorator.QueryHandler[retrieveUser, user.User]

type retrieveUserHandler struct {
	repo user.Repository
}

func NewRetrieveUser(uuid string) retrieveUser {
	return retrieveUser{uuid}
}

func NewRetrieveUserHandler(repo user.Repository, logger *logrus.Entry) RetrieveUserHandler {
	return decorator.ApplyQueryHandler(retrieveUserHandler{repo}, logger)
}

func (rh retrieveUserHandler) Handle(retrieveUser retrieveUser) (result user.User, err error) {
	result, err = rh.repo.Get(retrieveUser.uuid)
	if err != nil {
		return user.User{}, err
	}
	return result, nil
}
