package query

import (
	"backend/internal/common/decorator"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type RetrieveUser struct {
	Uuid string
}

type RetrieveUserHandler decorator.QueryHandler[RetrieveUser, user.User]

type retrieveUserHandler struct {
	userRepository user.Repository
}

func NewRetrieveUserHandler(repo user.Repository, logger *logrus.Entry) RetrieveUserHandler {
	return decorator.ApplyQueryHandler(retrieveUserHandler{userRepository: repo}, logger)
}

func (rh retrieveUserHandler) Handle(retrieveUser RetrieveUser) (result user.User, err error) {
	result, err = rh.userRepository.Get(retrieveUser.Uuid)
	if err != nil {
		return user.User{}, err
	}
	return result, nil
}
