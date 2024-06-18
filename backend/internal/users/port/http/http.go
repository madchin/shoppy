package ports

import (
	custom_error "backend/internal/common/errors"
	"backend/internal/common/server"
	"backend/internal/users/app"
	"backend/internal/users/app/command"
	"backend/internal/users/app/query"
	"backend/internal/users/domain/user"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type httpServer struct {
	app          app.Application
	authProvider server.Auth
}

func NewHttpServer(app app.Application, authProvider server.Auth) ServerInterface {
	return httpServer{app, authProvider}
}

func (h httpServer) PostUserLogin(w http.ResponseWriter, r *http.Request) {
	decodedUser, err := server.DecodeJSON[User](r.Body)
	if err != nil {
		server.ErrorHandler(w, r, custom_error.UnknownError("user login"))
		return
	}
	if decodedUser.Password == nil {
		server.ErrorHandler(w, r, custom_error.NewValidationError("user login", "password not provided"))
		return
	}

	user := query.NewLoginUser(decodedUser.Email, *decodedUser.Password)

	u, cerr := h.app.Query.LoginUser.Handle(r.Context(), user)
	if cerr.Error() != "" {
		server.ErrorHandler(w, r, cerr)
	}
	userInfo := server.NewUserInfo(u.Uuid())
	token, err := h.authProvider.Sign(userInfo)
	if err != nil {
		server.ErrorHandler(w, r, custom_error.UnknownError("user login"))
		return
	}
	w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", token))
	server.Success(w, http.StatusOK)
}

func (h httpServer) DeleteUserAddress(w http.ResponseWriter, r *http.Request, params DeleteUserAddressParams) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("delete user address", "unable to retrieve uuid"))
		return
	}
	deleteAddress := command.NewDeleteAddress(userInfo.Uuid, params.Street)
	err := h.app.Command.DeleteOneAddress.Handle(r.Context(), deleteAddress)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusNoContent)
}

func (h httpServer) DeleteUserAddresses(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("delete user addresses", "unable to retrieve uuid"))
		return
	}
	deleteAddresses := command.NewDeleteAllAddresses(userInfo.Uuid)
	err := h.app.Command.DeleteAllAddresses.Handle(r.Context(), deleteAddresses)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusNoContent)
}

func (h httpServer) GetUserAddress(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("retrieve user address", "unable to retrieve uuid"))
		return
	}
	retrieveAddress := query.NewRetrieveAddresses(userInfo.Uuid)
	addresses, err := h.app.Query.RetrieveAddresses.Handle(r.Context(), retrieveAddress)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	responseBody := mapDomainAddressesToHttpAddresses(addresses)
	server.SuccessWithBody(w, http.StatusOK, responseBody)
}

func (h httpServer) PostUserAddress(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("user address add", "unable to retrieve uuid"))
		return
	}
	address, err := server.DecodeJSON[Address](r.Body)
	if err != nil {
		server.ErrorHandler(w, r, custom_error.UnknownError("user address add"))
	}
	domainAddress := user.NewAddress(address.PostalCode, address.Street, address.Country, address.City)
	createAddress := command.NewCreateAddress(userInfo.Uuid, domainAddress)
	cerr := h.app.Command.CreateAddress.Handle(r.Context(), createAddress)
	if cerr.Error() != "" {
		server.ErrorHandler(w, r, cerr)
		return
	}
	server.SuccessWithBody(w, http.StatusCreated, Address{City: address.City, Country: address.Country, PostalCode: address.PostalCode, Street: address.Street})
}

func (h httpServer) PutUserAddress(w http.ResponseWriter, r *http.Request, params PutUserAddressParams) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("user address add", "unable to retrieve uuid"))
		return
	}
	address, err := server.DecodeJSON[Address](r.Body)
	if err != nil {
		server.ErrorHandler(w, r, custom_error.UnknownError("user address add"))
	}
	domainAddress := user.NewAddress(address.PostalCode, address.Street, address.Country, address.City)
	updateAddress := command.NewUpdateAddress(userInfo.Uuid, params.Street, domainAddress)
	cerr := h.app.Command.UpdateAddress.Handle(r.Context(), updateAddress)
	if cerr.Error() != "" {
		server.ErrorHandler(w, r, cerr)
		return
	}
	server.SuccessWithBody(w, http.StatusOK, Address{City: address.City, Country: address.Country, PostalCode: address.PostalCode, Street: address.Street})
}

func (h httpServer) GetUserPhones(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("retrieve user phones", "unable to retrieve uuid"))
		return
	}
	retrievePhones := query.NewRetrievePhones(userInfo.Uuid)
	phones, err := h.app.Query.RetrievePhones.Handle(r.Context(), retrievePhones)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	numbers := phones.AllPhoneNumbers()
	server.SuccessWithBody(w, http.StatusCreated, Phones{Numbers: numbers})
}

func (h httpServer) DeleteUserPhone(w http.ResponseWriter, r *http.Request, params DeleteUserPhoneParams) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("delete user phone", "unable to retrieve uuid"))
		return
	}
	deletePhone := command.NewDeletePhone(userInfo.Uuid, user.NewPhone(params.Number))
	err := h.app.Command.DeleteOnePhone.Handle(r.Context(), deletePhone)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusNoContent)
}

func (h httpServer) DeleteUserPhones(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("delete user phones", "unable to retrieve uuid"))
		return
	}
	deleteAllPhones := command.NewDeleteAllPhones(userInfo.Uuid)
	err := h.app.Command.DeleteAllPhones.Handle(r.Context(), deleteAllPhones)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusNoContent)
}

func (h httpServer) PostUserPhone(w http.ResponseWriter, r *http.Request, params PostUserPhoneParams) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("create user phone", "unable to retrieve uuid"))
		return
	}
	domainPhone := user.NewPhone(params.Number)
	createPhone := command.NewCreatePhone(userInfo.Uuid, domainPhone)
	if err := h.app.Command.CreatePhone.Handle(r.Context(), createPhone); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	number := domainPhone.Number()
	server.SuccessWithBody(w, http.StatusCreated, Phone{Number: number})
}

func (h httpServer) PutUserPhone(w http.ResponseWriter, r *http.Request, params PutUserPhoneParams) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("user phone add", "unable to retrieve uuid"))
		return
	}
	phone, err := server.DecodeJSON[Phone](r.Body)
	if err != nil {
		server.ErrorHandler(w, r, custom_error.UnknownError("user phone add"))
		return
	}
	nextPhone := user.NewPhone(phone.Number)
	updatePhone := command.NewUpdatePhone(userInfo.Uuid, params.PreviousNumber, nextPhone)
	if err := h.app.Command.UpdatePhone.Handle(r.Context(), updatePhone); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	number := nextPhone.Number()
	server.SuccessWithBody(w, http.StatusOK, Phone{Number: number})
}

func (h httpServer) DeleteUserDetail(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("delete user detail", fmt.Sprintf("unable to retrieve uuid: user info is %s", userInfo.Uuid)))
		return
	}
	user := command.NewDeleteUserDetail(userInfo.Uuid)
	err := h.app.Command.DeleteUserDetail.Handle(r.Context(), user)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusNoContent)
}

func (h httpServer) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("retrieve user details", "unable to retrieve uuid"))
		return
	}
	user := query.NewRetrieveUserDetail(userInfo.Uuid)
	u, err := h.app.Query.RetrieveUserDetail.Handle(r.Context(), user)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	firstName, lastName := u.FirstName(), u.LastName()
	server.SuccessWithBody(w, http.StatusOK, UserDetail{FirstName: firstName, LastName: lastName})
}

func (h httpServer) PostUserDetail(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("user detail add", "unable to retrieve uuid"))
		return
	}
	decodedUser, err := server.DecodeJSON[UserDetail](r.Body)
	if err != nil {
		server.ErrorHandler(w, r, custom_error.UnknownError("user detail add"))
		return
	}
	domainUser := user.NewUserDetail(decodedUser.FirstName, decodedUser.LastName)
	createUserDetail := command.NewCreateUserDetail(userInfo.Uuid, domainUser)
	if err := h.app.Command.CreateUserDetail.Handle(r.Context(), createUserDetail); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	firstName, lastName := domainUser.FirstName(), domainUser.LastName()
	server.SuccessWithBody(w, http.StatusCreated, UserDetail{FirstName: firstName, LastName: lastName})
}

func (h httpServer) PutUserDetailUpdateFirstName(w http.ResponseWriter, r *http.Request, params PutUserDetailUpdateFirstNameParams) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("update user detail first name", "unable to retrieve uuid"))
		return
	}
	updateUserDetailFirstName := command.NewUpdateUserDetailFirstName(userInfo.Uuid, params.FirstName)
	if err := h.app.Command.UpdateUserDetailFirstName.Handle(r.Context(), updateUserDetailFirstName); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.SuccessWithBody(w, http.StatusOK, UserDetail{FirstName: params.FirstName})
}

func (h httpServer) PutUserDetailUpdateLastName(w http.ResponseWriter, r *http.Request, params PutUserDetailUpdateLastNameParams) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("update user detail last name", "unable to retrieve uuid"))
		return
	}
	updateUserDetailLastName := command.NewUpdateUserDetailLastName(userInfo.Uuid, params.LastName)
	if err := h.app.Command.UpdateUserDetailLastName.Handle(r.Context(), updateUserDetailLastName); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.SuccessWithBody(w, http.StatusOK, UserDetail{LastName: params.LastName})
}

func (h httpServer) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("delete user", "unable to retrieve uuid"))
		return
	}
	user := command.NewDeleteUser(userInfo.Uuid)
	err := h.app.Command.DeleteUser.Handle(r.Context(), user)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusNoContent)
}

func (h httpServer) GetUser(w http.ResponseWriter, r *http.Request) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("retrieve user", "unable to retrieve uuid"))
		return
	}
	user := query.NewRetrieveUser(userInfo.Uuid)
	u, err := h.app.Query.RetrieveUser.Handle(r.Context(), user)
	if err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.SuccessWithBody(w, http.StatusOK, User{Email: u.Email(), Name: u.Name()})
}

func (h httpServer) PostUser(w http.ResponseWriter, r *http.Request) {
	uuid := uuid.NewString()
	decodedUser, err := server.DecodeJSON[User](r.Body)
	if err != nil {
		server.ErrorHandler(w, r, custom_error.UnknownError("user add"))
		return
	}
	if decodedUser.Password == nil {
		server.ErrorHandler(w, r, custom_error.NewValidationError("user add", "password has not been provided"))
		return
	}
	domainUser := user.NewUser(uuid, *decodedUser.Password, decodedUser.Name, decodedUser.Email)
	registerUser := command.NewRegisterUser(uuid, domainUser)
	if err := h.app.Command.RegisterUser.Handle(r.Context(), registerUser); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	email, name, password := domainUser.Email(), domainUser.Name(), domainUser.Password()
	server.SuccessWithBody(w, http.StatusCreated, User{Email: email, Name: name, Password: &password})
}

func (h httpServer) PutUserUpdateEmail(w http.ResponseWriter, r *http.Request, params PutUserUpdateEmailParams) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("update user email", "unable to retrieve uuid"))
		return
	}
	updateUserEmail := command.NewUpdateUserEmail(userInfo.Uuid, params.Email)
	if err := h.app.Command.UpdateUserEmail.Handle(r.Context(), updateUserEmail); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.SuccessWithBody(w, http.StatusOK, User{Email: params.Email})
}

func (h httpServer) PutUserUpdateName(w http.ResponseWriter, r *http.Request, params PutUserUpdateNameParams) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("update user name", "unable to retrieve uuid"))
		return
	}
	updateUserName := command.NewUpdateUserName(userInfo.Uuid, params.Name)
	if err := h.app.Command.UpdateUserName.Handle(r.Context(), updateUserName); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.SuccessWithBody(w, http.StatusOK, User{Name: params.Name})
}

func (h httpServer) PutUserUpdatePassword(w http.ResponseWriter, r *http.Request, params PutUserUpdatePasswordParams) {
	userInfo, ok := server.UserInfoFromContext(r.Context())
	if !ok {
		server.ErrorHandler(w, r, custom_error.NewAuthorizationError("update user password", "unable to retrieve uuid"))
		return
	}
	updateUserPassword := command.NewUpdateUserPassword(userInfo.Uuid, params.Password)
	if err := h.app.Command.UpdateUserPassword.Handle(r.Context(), updateUserPassword); err.Error() != "" {
		server.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusOK)
}
