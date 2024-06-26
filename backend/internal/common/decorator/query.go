package decorator

import (
	custom_error "backend/internal/common/errors"
	"context"

	"github.com/sirupsen/logrus"
)

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, cmd Q) (result R, err custom_error.ContextError)
}

func ApplyQueryHandler[Q any, R any](base QueryHandler[Q, R], logger *logrus.Entry) QueryHandler[Q, R] {
	return queryHandler[Q, R]{
		base:   base,
		logger: logger,
	}
}
