package app

import (
	"backend/internal/users/app/command"
	"backend/internal/users/app/query"
	"backend/internal/users/domain/user"

	"github.com/sirupsen/logrus"
)

//usecases

type Command struct {
	RegisterUser              command.RegisterUserHandler
	DeleteUser                command.DeleteUserHandler
	UpdateUserEmail           command.UpdateUserEmailHandler
	UpdateUserName            command.UpdateUserNameHandler
	UpdateUserPassword        command.UpdateUserPasswordHandler
	DeleteUserDetail          command.DeleteUserDetailHandler
	CreateUserDetail          command.CreateUserDetailHandler
	UpdateUserDetailFirstName command.UpdateUserDetailFirstNameHandler
	UpdateUserDetailLastName  command.UpdateUserDetailLastNameHandler
}

type Query struct {
	RetrieveUser       query.RetrieveUserHandler
	RetrieveUserDetail query.RetrieveUserDetailHandler
}

type Application struct {
	Command Command
	Query   Query
}

func NewApplication(userRepository user.Repository, logger *logrus.Entry) Application {
	return Application{
		Command: Command{
			RegisterUser:       command.NewRegisterUserHandler(userRepository, logger),
			DeleteUser:         command.NewDeleteUserHandler(userRepository, logger),
			UpdateUserEmail:    command.NewUpdateUserEmailHandler(userRepository, logger),
			UpdateUserName:     command.NewUpdateUserNameHandler(userRepository, logger),
			UpdateUserPassword: command.NewUpdateUserPasswordHandler(userRepository, logger),
		},
		Query: Query{RetrieveUser: query.NewRetrieveUserHandler(userRepository, logger)},
	}
}
