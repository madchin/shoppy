package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"context"

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

func NewUpdatePhone(userUuid string, prevNumber string, phone user.Phone) updatePhone {
	return updatePhone{userUuid, prevNumber, phone}
}

func NewUpdatePhoneHandler(pr user.PhoneRepository, logger *logrus.Entry) UpdatePhoneHandler {
	return decorator.ApplyCommandHandler(updatePhoneHandler{pr}, logger)
}

func (uph updatePhoneHandler) Handle(ctx context.Context, cmd updatePhone) custom_error.ContextError {
	return uph.pr.Update(ctx, cmd.userUuid, cmd.prevNumber, cmd.phone, func(p user.Phone) []error {
		return p.Validate()
	})
}
