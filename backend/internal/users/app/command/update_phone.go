package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type updatePhone struct {
	userUuid   string
	prevNumber string
	phone      user.Phone
}

type updatePhoneHandler struct {
	pr user.PhoneRepository
}

type UpdatePhoneHandler decorator.CommandHandler[updatePhone]

func NewUpdatePhoneHandler(pr user.PhoneRepository, logger *logrus.Entry) UpdatePhoneHandler {
	return decorator.ApplyCommandHandler(updatePhoneHandler{pr}, logger)
}

func (uph updatePhoneHandler) Handle(cmd updatePhone) custom_error.ContextError {
	return uph.pr.Update(cmd.userUuid, cmd.prevNumber, cmd.phone, func(p user.Phone) []error {
		return p.Validate()
	})
}
