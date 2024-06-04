package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type deleteAllPhones struct {
	userUuid string
}

type deleteAllPhonesHandler struct {
	pr user.PhoneRepository
}

type DeleteAllPhonesHandler decorator.CommandHandler[deleteAllPhones]

func NewDeleteAllPhonesHandler(pr user.PhoneRepository, logger *logrus.Entry) DeleteAllPhonesHandler {
	return decorator.ApplyCommandHandler(deleteAllPhonesHandler{pr}, logger)
}

func (daph deleteAllPhonesHandler) Handle(cmd deleteAllPhones) custom_error.ContextError {
	return daph.pr.DeleteAll(cmd.userUuid)
}
