package query

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type retrieveAddress struct {
	userUuid string
}

type retrieveAddressHandler struct {
	repo user.AddressRepository
}

type RetrieveAddressHandler decorator.QueryHandler[retrieveAddress, user.Address]

func NewRetrieveAddress(userUuid string) retrieveAddress {
	return retrieveAddress{userUuid}
}

func NewRetrieveAddressHandler(ar user.AddressRepository, logger *logrus.Entry) RetrieveAddressHandler {
	return decorator.ApplyQueryHandler(retrieveAddressHandler{ar}, logger)
}

func (rah retrieveAddressHandler) Handle(q retrieveAddress) (user.Address, custom_error.ContextError) {
	return rah.repo.Get(q.userUuid)
}
