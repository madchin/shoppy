package query

import (
	"backend/internal/common/decorator"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"context"

	"github.com/sirupsen/logrus"
)

type retrievePhones struct {
	userUuid string
}

type retrievePhonesHandler struct {
	pr user.PhoneRepository
}

type RetrievePhonesHandler decorator.QueryHandler[retrievePhones, user.Phones]

func NewRetrievePhones(userUuid string) retrievePhones {
	return retrievePhones{userUuid}
}

func NewRetrievePhonesHandler(pr user.PhoneRepository, logger *logrus.Entry) RetrievePhonesHandler {
	return decorator.ApplyQueryHandler(retrievePhonesHandler{pr}, logger)
}

func (rph retrievePhonesHandler) Handle(ctx context.Context, q retrievePhones) (user.Phones, custom_error.ContextError) {
	return rph.pr.Get(ctx, q.userUuid)
}
