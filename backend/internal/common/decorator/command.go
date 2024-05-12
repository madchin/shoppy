package decorator

import "github.com/sirupsen/logrus"

type CommandHandler[C any] interface {
	Handle(cmd C) error
}

func ApplyCommandHandler[C any](base CommandHandler[C], logger *logrus.Entry) CommandHandler[C] {
	return commandHandler[C]{
		base:   base,
		logger: logger,
	}
}
