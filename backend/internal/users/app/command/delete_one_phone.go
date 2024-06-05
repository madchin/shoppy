package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

type deletePhone struct {
	userUuid string
	phone    user.Phone
}

type deletePhoneHandler struct {
	pr user.PhoneRepository
}

type DeleteOnePhoneHandler decorator.CommandHandler[deletePhone]

func NewDeletePhone(userUuid string, phone user.Phone) deletePhone {
	return deletePhone{userUuid, phone}
}

func NewDeleteOnePhoneHandler(pr user.PhoneRepository, logger *logrus.Entry) DeleteOnePhoneHandler {
	return decorator.ApplyCommandHandler(deletePhoneHandler{pr}, logger)
}

func (dph deletePhoneHandler) Handle(cmd deletePhone) custom_error.ContextError {
	return dph.pr.DeletePhone(cmd.userUuid, cmd.phone, func(phones user.Phones) error {
		if !phones.NumberExist(cmd.phone.Number()) {
			return custom_error.NewValidationError("phone deletion", "provided phone do not exist")
		}
		return custom_error.ContextError{}
	})
}
