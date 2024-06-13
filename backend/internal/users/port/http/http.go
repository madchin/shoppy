package ports

import (
	custom_error "backend/internal/common/errors"
	"backend/internal/common/server"
	"backend/internal/users/app"
	"backend/internal/users/app/command"
	"backend/internal/users/app/query"
	"backend/internal/users/domain/user"
	"net/http"
)

type httpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) ServerInterface {
	return httpServer{app}
}

func (h httpServer) DeleteUserAddress(w http.ResponseWriter, r *http.Request, params DeleteUserAddressParams) {
	uuid := ""
	deleteAddress := command.NewDeleteAddress(uuid, params.Street)
	err := h.app.Command.DeleteOneAddress.Handle(deleteAddress)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusNoContent)
}

func (h httpServer) DeleteUserAddresses(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	deleteAddresses := command.NewDeleteAllAddresses(uuid)
	err := h.app.Command.DeleteAllAddresses.Handle(deleteAddresses)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusNoContent)
}

func (h httpServer) GetUserAddress(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	retrieveAddress := query.NewRetrieveAddresses(uuid)
	addresses, err := h.app.Query.RetrieveAddresses.Handle(retrieveAddress)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	responseBody := mapDomainAddressesToResponseAddresses(addresses)
	server.SuccessWithBody(w, http.StatusOK, responseBody)
}

func (h httpServer) PostUserAddress(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	address, err := server.DecodeJSON[Address](r.Body)
	if err != nil {
		server.ErrorHandler(w, r, custom_error.UnknownError("user address add"))
	}
	domainAddress := user.NewAddress(address.PostalCode, address.Street, address.Country, address.City)
	createAddress := command.NewCreateAddress(uuid, domainAddress)
	cerr := h.app.Command.CreateAddress.Handle(createAddress)
	if cerr.Error() != "" {
		server.ErrorHandler(w, r, cerr)
		return
	}
	server.SuccessWithBody(w, http.StatusCreated, Address{City: address.City, Country: address.Country, PostalCode: address.PostalCode, Street: address.Street})
}

func (h httpServer) PutUserAddress(w http.ResponseWriter, r *http.Request, params PutUserAddressParams) {
	uuid := ""
	address, err := server.DecodeJSON[Address](r.Body)
	if err != nil {
		server.ErrorHandler(w, r, custom_error.UnknownError("user address add"))
	}
	domainAddress := user.NewAddress(address.PostalCode, address.Street, address.Country, address.City)
	updateAddress := command.NewUpdateAddress(uuid, params.Street, domainAddress)
	cerr := h.app.Command.UpdateAddress.Handle(updateAddress)
	if cerr.Error() != "" {
		server.ErrorHandler(w, r, cerr)
		return
	}
	server.SuccessWithBody(w, http.StatusOK, Address{City: address.City, Country: address.Country, PostalCode: address.PostalCode, Street: address.Street})
}

func (h httpServer) GetUserPhones(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	retrievePhones := query.NewRetrievePhones(uuid)
	phones, err := h.app.Query.RetrievePhones.Handle(retrievePhones)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	numbers := phones.AllPhoneNumbers()
	server.SuccessWithBody(w, http.StatusCreated, Phones{Numbers: numbers})
}

func (h httpServer) DeleteUserPhone(w http.ResponseWriter, r *http.Request, params DeleteUserPhoneParams) {
	uuid := ""
	deletePhone := command.NewDeletePhone(uuid, user.NewPhone(params.Number))
	err := h.app.Command.DeleteOnePhone.Handle(deletePhone)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusNoContent)
}

func (h httpServer) DeleteUserPhones(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	deleteAllPhones := command.NewDeleteAllPhones(uuid)
	err := h.app.Command.DeleteAllPhones.Handle(deleteAllPhones)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusNoContent)
}

func (h httpServer) PostUserPhone(w http.ResponseWriter, r *http.Request, params PostUserPhoneParams) {
	uuid := ""
	domainPhone := user.NewPhone(params.Number)
	createPhone := command.NewCreatePhone(uuid, domainPhone)
	if err := h.app.Command.CreatePhone.Handle(createPhone); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	number := domainPhone.Number()
	server.SuccessWithBody(w, http.StatusCreated, Phone{Number: number})
}

func (h httpServer) PutUserPhone(w http.ResponseWriter, r *http.Request, params PutUserPhoneParams) {
	uuid := ""
	nextPhone, err := server.DecodeJSON[user.Phone](r.Body)
	if err != nil {
		server.ErrorHandler(w, r, custom_error.UnknownError("user phone add"))
		return
	}
	updatePhone := command.NewUpdatePhone(uuid, params.PreviousNumber, *nextPhone)
	if err := h.app.Command.UpdatePhone.Handle(updatePhone); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	number := nextPhone.Number()
	server.SuccessWithBody(w, http.StatusOK, Phone{Number: number})
}

func (h httpServer) DeleteUserDetail(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	user := command.NewDeleteUserDetail(uuid)
	err := h.app.Command.DeleteUserDetail.Handle(user)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusNoContent)
}

func (h httpServer) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	user := query.NewRetrieveUserDetail(uuid)
	u, err := h.app.Query.RetrieveUserDetail.Handle(user)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	firstName, lastName := u.FirstName(), u.LastName()
	server.SuccessWithBody(w, http.StatusOK, UserDetail{FirstName: firstName, LastName: lastName})
}

func (h httpServer) PostUserDetail(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	decodedUser, err := server.DecodeJSON[UserDetail](r.Body)
	if err != nil {
		server.ErrorHandler(w, r, custom_error.UnknownError("user detail add"))
		return
	}
	domainUser := user.NewUserDetail(decodedUser.FirstName, decodedUser.LastName)
	createUserDetail := command.NewCreateUserDetail(uuid, domainUser)
	if err := h.app.Command.CreateUserDetail.Handle(createUserDetail); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	firstName, lastName := domainUser.FirstName(), domainUser.LastName()
	server.SuccessWithBody(w, http.StatusCreated, UserDetail{FirstName: firstName, LastName: lastName})
}

func (h httpServer) PutUserDetailUpdateFirstName(w http.ResponseWriter, r *http.Request, params PutUserDetailUpdateFirstNameParams) {
	uuid := ""
	updateUserDetailFirstName := command.NewUpdateUserDetailFirstName(uuid, params.FirstName)
	if err := h.app.Command.UpdateUserDetailFirstName.Handle(updateUserDetailFirstName); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.SuccessWithBody(w, http.StatusOK, UserDetail{FirstName: params.FirstName})
}

func (h httpServer) PutUserDetailUpdateLastName(w http.ResponseWriter, r *http.Request, params PutUserDetailUpdateLastNameParams) {
	uuid := ""
	updateUserDetailLastName := command.NewUpdateUserDetailLastName(uuid, params.LastName)
	if err := h.app.Command.UpdateUserDetailLastName.Handle(updateUserDetailLastName); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.SuccessWithBody(w, http.StatusOK, UserDetail{LastName: params.LastName})
}

func (h httpServer) DeleteUser(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	user := command.NewDeleteUser(uuid)
	err := h.app.Command.DeleteUser.Handle(user)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusNoContent)
}

func (h httpServer) GetUser(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	user := query.NewRetrieveUser(uuid)
	u, err := h.app.Query.RetrieveUser.Handle(user)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.SuccessWithBody(w, http.StatusOK, User{Email: u.Email(), Name: u.Name()})
}

func (h httpServer) PostUser(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	decodedUser, err := server.DecodeJSON[User](r.Body)
	if err != nil {
		server.ErrorHandler(w, r, custom_error.UnknownError("user add"))
		return
	}
	domainUser := user.NewUser(decodedUser.Name, decodedUser.Email, *decodedUser.Password)
	registerUser := command.NewRegisterUser(uuid, domainUser)
	if err := h.app.Command.RegisterUser.Handle(registerUser); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	email, name, password := domainUser.Email(), domainUser.Name(), domainUser.Password()
	server.SuccessWithBody(w, http.StatusCreated, User{Email: email, Name: name, Password: &password})
}

func (h httpServer) PutUserUpdateEmail(w http.ResponseWriter, r *http.Request, params PutUserUpdateEmailParams) {
	uuid := ""
	updateUserEmail := command.NewUpdateUserEmail(uuid, params.Email)
	if err := h.app.Command.UpdateUserEmail.Handle(updateUserEmail); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.SuccessWithBody(w, http.StatusOK, User{Email: params.Email})
}

func (h httpServer) PutUserUpdateName(w http.ResponseWriter, r *http.Request, params PutUserUpdateNameParams) {
	uuid := ""
	updateUserName := command.NewUpdateUserName(uuid, params.Name)
	if err := h.app.Command.UpdateUserName.Handle(updateUserName); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.SuccessWithBody(w, http.StatusOK, User{Name: params.Name})
}

func (h httpServer) PutUserUpdatePassword(w http.ResponseWriter, r *http.Request, params PutUserUpdatePasswordParams) {
	uuid := ""
	updateUserPassword := command.NewUpdateUserPassword(uuid, params.Password)
	if err := h.app.Command.UpdateUserPassword.Handle(updateUserPassword); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusOK)
}
