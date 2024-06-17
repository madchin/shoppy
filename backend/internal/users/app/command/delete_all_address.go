package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"context"

	"github.com/sirupsen/logrus"
)

type deleteAllAddresses struct {
	userUuid string
}

type deleteAllAddressesHandler struct {
	repo user.AddressRepository
}

type DeleteAllAddressesHandler decorator.CommandHandler[deleteAllAddresses]

func NewDeleteAllAddresses(userUuid string) deleteAllAddresses {
	return deleteAllAddresses{userUuid}
}

func NewDeleteAllAddressesHandler(repo user.AddressRepository, logger *logrus.Entry) DeleteAllAddressesHandler {
	return decorator.ApplyCommandHandler(deleteAllAddressesHandler{repo}, logger)
}

func (dah deleteAllAddressesHandler) Handle(ctx context.Context, cmd deleteAllAddresses) custom_error.ContextError {
	return dah.repo.DeleteAll(ctx, cmd.userUuid)
}
