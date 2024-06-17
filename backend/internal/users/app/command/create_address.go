package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"context"

	"github.com/sirupsen/logrus"
)

type createAddress struct {
	userUuid string
	address  user.Address
}

type createAddressHandler struct {
	repo user.AddressRepository
}

type CreateAddressHandler decorator.CommandHandler[createAddress]

func NewCreateAddress(userId string, address user.Address) createAddress {
	return createAddress{userId, address}
}

func NewCreateAddressHandler(addressRepository user.AddressRepository, logger *logrus.Entry) CreateAddressHandler {
	return decorator.ApplyCommandHandler(createAddressHandler{addressRepository}, logger)
}

func (c createAddressHandler) Handle(ctx context.Context, cmd createAddress) custom_error.ContextError {
	return c.repo.Create(ctx, cmd.userUuid, cmd.address, func(a user.Address) []error {
		return a.Validate()
	})
}
