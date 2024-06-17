package command

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"context"

	"github.com/sirupsen/logrus"
)

type createPhone struct {
	userUuid string
	phone    user.Phone
}

type createPhoneHandler struct {
	pr user.PhoneRepository
}

type CreatePhoneHandler decorator.CommandHandler[createPhone]

func NewCreatePhone(userUuid string, phone user.Phone) createPhone {
	return createPhone{userUuid, phone}
}

func NewCreatePhoneHandler(phoneRepository user.PhoneRepository, logger *logrus.Entry) CreatePhoneHandler {
	return decorator.ApplyCommandHandler(createPhoneHandler{phoneRepository}, logger)
}

func (cph createPhoneHandler) Handle(ctx context.Context, cmd createPhone) custom_error.ContextError {
	return cph.pr.Create(ctx, cmd.userUuid, cmd.phone, func(p user.Phone) []error {
		return p.Validate()
	})
}
