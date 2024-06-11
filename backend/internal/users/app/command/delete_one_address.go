package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type deleteOneAddress struct {
	userUuid string
	street   string
}

type deleteOneAddressHandler struct {
	repo user.AddressRepository
}

type DeleteAddressHandler decorator.CommandHandler[deleteOneAddress]

func NewDeleteAddress(userUuid string, street string) deleteOneAddress {
	return deleteOneAddress{userUuid, street}
}

func NewDeleteAddressHandler(repo user.AddressRepository, logger *logrus.Entry) DeleteAddressHandler {
	return decorator.ApplyCommandHandler(deleteOneAddressHandler{repo}, logger)
}

func (dah deleteOneAddressHandler) Handle(cmd deleteOneAddress) custom_error.ContextError {
	return dah.repo.DeleteAddress(cmd.userUuid, cmd.street)
}