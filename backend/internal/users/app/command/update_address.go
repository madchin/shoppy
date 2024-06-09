package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type updateAddress struct {
	userUuid   string
	prevStreet string
	address    user.Address
}

type updateAddressHandler struct {
	repo user.AddressRepository
}

type UpdateAddressHandler decorator.CommandHandler[updateAddress]

func NewUpdateAddress(userUuid string, prevStreet string, address user.Address) updateAddress {
	return updateAddress{userUuid, prevStreet, address}
}

func NewUpdateAddressHandler(repo user.AddressRepository, logger *logrus.Entry) UpdateAddressHandler {
	return decorator.ApplyCommandHandler(updateAddressHandler{repo}, logger)
}

func (c updateAddressHandler) Handle(cmd updateAddress) custom_error.ContextError {
	return c.repo.Update(cmd.userUuid, cmd.prevStreet, cmd.address, func(a user.Address) []error {
		return a.Validate()
	})
}
