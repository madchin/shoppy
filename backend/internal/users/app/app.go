package app

import (
	"backend/internal/users/app/query"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

//usecases

type Command struct {
	CreateUser func()
}

type Query struct {
	RetrieveUser query.RetrieveUserHandler
}

type Application struct {
	Command Command
	Query   Query
}

func NewApplication(userRepository user.Repository, logger *logrus.Entry) Application {
	return Application{
		Command: Command{},
		Query:   Query{RetrieveUser: query.NewRetrieveUserHandler(userRepository, logger)},
	}
}
