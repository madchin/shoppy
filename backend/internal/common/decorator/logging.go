package decorator

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type queryHandler[Q any, R any] struct {
	base   QueryHandler[Q, R]
	logger *logrus.Entry
}

func (q queryHandler[Q, R]) Handle(cmd Q) (result R, err error) {
	logger := q.logger.WithFields(logrus.Fields{
		"query_name": fmt.Sprintf("%T", cmd),
		"query_body": fmt.Sprintf("%v", cmd),
	})

	logger.Debug("Executing query")

	defer func() {
		if err != nil {
			logger.WithError(err).Error("Query execution failed")
			return
		}
		logger.Info("Query executed successfully")
	}()

	return q.base.Handle(cmd)
}

type commandHandler[C any] struct {
	base   CommandHandler[C]
	logger *logrus.Entry
}

func (c commandHandler[C]) Handle(cmd C) (err error) {
	logger := c.logger.WithFields(logrus.Fields{
		"command_name": fmt.Sprintf("%T", cmd),
		"command_body": fmt.Sprintf("%v", cmd),
	})
	logger.Debug("Executing command")

	defer func() {
		if err != nil {
			logger.WithError(err).Error("Command execution failed")
			return
		}
		logger.Info("Command executed successfully")
	}()

	return c.base.Handle(cmd)
}
