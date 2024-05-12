package decorator

import "github.com/sirupsen/logrus"

type QueryHandler[Q any, R any] interface {
	Handle(cmd Q) (result R, err error)
}

func ApplyQueryHandler[Q any, R any](base QueryHandler[Q, R], logger *logrus.Entry) QueryHandler[Q, R] {
	return queryHandler[Q, R]{
		base:   base,
		logger: logger,
	}
}
