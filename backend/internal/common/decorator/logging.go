package decorator

import (
	custom_error "backend/internal/common/errors"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type queryHandler[Q any, R any] struct {
	base   QueryHandler[Q, R]
	logger *logrus.Entry
}

func (q queryHandler[Q, R]) Handle(ctx context.Context, cmd Q) (result R, err custom_error.ContextError) {
	logger := q.logger.WithFields(logrus.Fields{
		"query_name": fmt.Sprintf(" %T", cmd),
		"query_body": fmt.Sprintf(" %v", cmd),
	})

	logger.Debug("Executing query")

	defer func() {
		if err.Error() != "" {
			logger.WithField(err.Type().String(), err.Error()).Error("Query execution failed")
			return
		}
		logger.Info("Query executed successfully")
	}()

	return q.base.Handle(ctx, cmd)
}

type commandHandler[C any] struct {
	base   CommandHandler[C]
	logger *logrus.Entry
}

func (c commandHandler[C]) Handle(ctx context.Context, cmd C) (err custom_error.ContextError) {
	logger := c.logger.WithFields(logrus.Fields{
		"command_name": fmt.Sprintf(" %T", cmd),
		"command_body": fmt.Sprintf(" %v", cmd),
	})
	logger.Debug("Executing command")

	defer func() {
		if err.Error() != "" {
			logger.WithField(err.Type().String(), err.Error()).Error("Command execution failed")
			return
		}
		logger.Info("Command executed successfully")
	}()

	return c.base.Handle(ctx, cmd)
}
