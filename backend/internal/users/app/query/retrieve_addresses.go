package query

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type retrieveAddresses struct {
	userUuid string
}

type retrieveAddressesHandler struct {
	repo user.AddressRepository
}

type RetrieveAddressHandler decorator.QueryHandler[retrieveAddresses, user.Addresses]

func NewRetrieveAddress(userUuid string) retrieveAddresses {
	return retrieveAddresses{userUuid}
}

func NewRetrieveAddressHandler(ar user.AddressRepository, logger *logrus.Entry) RetrieveAddressHandler {
	return decorator.ApplyQueryHandler(retrieveAddressesHandler{ar}, logger)
}

func (rah retrieveAddressesHandler) Handle(q retrieveAddresses) (user.Addresses, custom_error.ContextError) {
	return rah.repo.Get(q.userUuid)
}
