package query

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type retrieveUserDetail struct {
	uuid string
}

type RetrieveUserDetailHandler decorator.QueryHandler[retrieveUserDetail, user.UserDetail]

type retrieveUserDetailHandler struct {
	repo user.DetailRepository
}

func NewRetrieveUserDetail(uuid string) retrieveUserDetail {
	return retrieveUserDetail{uuid}
}

func NewRetrieveUserDetailHandler(repo user.DetailRepository, logger *logrus.Entry) RetrieveUserDetailHandler {
	return decorator.ApplyQueryHandler(retrieveUserDetailHandler{repo}, logger)
}

func (rh retrieveUserDetailHandler) Handle(retrieveUser retrieveUserDetail) (user.UserDetail, custom_error.ContextError) {
	return rh.repo.Get(retrieveUser.uuid)
}
