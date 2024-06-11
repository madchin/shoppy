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
	CreatePhone               command.CreatePhoneHandler
	UpdatePhone               command.UpdatePhoneHandler
	DeleteOnePhone            command.DeleteOnePhoneHandler
	DeleteAllPhones           command.DeleteAllPhonesHandler
	CreateAddress             command.CreateAddressHandler
	DeleteOneAddress          command.DeleteOneAddressHandler
	DeleteAllAddresses        command.DeleteAllAddressesHandler
	UpdateAddress             command.UpdateAddressHandler
}

type Query struct {
	RetrieveUser       query.RetrieveUserHandler
	RetrieveUserDetail query.RetrieveUserDetailHandler
	RetrievePhones     query.RetrievePhonesHandler
	RetrieveAddresses  query.RetrieveAddressesHandler
}

type Application struct {
	Command Command
	Query   Query
}

func NewApplication(
	userRepository user.Repository,
	userDetailRepository user.DetailRepository,
	phoneRepository user.PhoneRepository,
	addressRepository user.AddressRepository,
	logger *logrus.Entry,
) Application {
	return Application{
		Command: Command{
			RegisterUser:              command.NewRegisterUserHandler(userRepository, logger),
			DeleteUser:                command.NewDeleteUserHandler(userRepository, logger),
			UpdateUserEmail:           command.NewUpdateUserEmailHandler(userRepository, logger),
			UpdateUserName:            command.NewUpdateUserNameHandler(userRepository, logger),
			UpdateUserPassword:        command.NewUpdateUserPasswordHandler(userRepository, logger),
			DeleteUserDetail:          command.NewDeleteUserDetailHandler(userDetailRepository, logger),
			CreateUserDetail:          command.NewCreateUserDetailHandler(userDetailRepository, logger),
			UpdateUserDetailFirstName: command.NewUpdateUserDetailFirstNameHandler(userDetailRepository, logger),
			UpdateUserDetailLastName:  command.NewUpdateUserDetailLastNameHandler(userDetailRepository, logger),
			CreatePhone:               command.NewCreatePhoneHandler(phoneRepository, logger),
			UpdatePhone:               command.NewUpdatePhoneHandler(phoneRepository, logger),
			DeleteOnePhone:            command.NewDeleteOnePhoneHandler(phoneRepository, logger),
			DeleteAllPhones:           command.NewDeleteAllPhonesHandler(phoneRepository, logger),
			CreateAddress:             command.NewCreateAddressHandler(addressRepository, logger),
			DeleteAllAddresses:        command.NewDeleteAllAddressesHandler(addressRepository, logger),
			DeleteOneAddress:          command.NewDeleteOneAddressHandler(addressRepository, logger),
			UpdateAddress:             command.NewUpdateAddressHandler(addressRepository, logger),
		},
		Query: Query{
			RetrieveUser:       query.NewRetrieveUserHandler(userRepository, logger),
			RetrieveUserDetail: query.NewRetrieveUserDetailHandler(userDetailRepository, logger),
			RetrievePhones:     query.NewRetrievePhonesHandler(phoneRepository, logger),
			RetrieveAddresses:  query.NewRetrieveAddressesHandler(addressRepository, logger),
		},
	}
}
