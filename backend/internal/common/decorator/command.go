package decorator

import (
	custom_error "backend/internal/common/errors"
	"context"

	"github.com/sirupsen/logrus"
)

type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) custom_error.ContextError
}

func ApplyCommandHandler[C any](base CommandHandler[C], logger *logrus.Entry) CommandHandler[C] {
	return commandHandler[C]{
		base:   base,
		logger: logger,
	}
}
