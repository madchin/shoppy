package ports

import (
	custom_error "backend/internal/common/errors"
	"backend/internal/common/server"
	"backend/internal/users/app"
	"backend/internal/users/app/command"
	"backend/internal/users/app/query"
	"backend/internal/users/domain/user"
	"backend/internal/users/port/http/httperror"
	"errors"
	"net/http"
)

type httpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) ServerInterface {
	return httpServer{app}
}

func (h httpServer) DeleteUser(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	user := command.NewDeleteUser(uuid)
	err := h.app.Command.DeleteUser.Handle(user)
	if err.Error() != "" {
		httperror.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusOK)
}

func (h httpServer) GetUser(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	user := query.NewRetrieveUser(uuid)
	u, err := h.app.Query.RetrieveUser.Handle(user)
	if err.Error() != "" {
		httperror.ErrorHandler(w, r, err)
		return
	}
	server.SuccessWithBody(w, http.StatusCreated, User{Email: u.Email(), Name: u.Name()})
}

func (h httpServer) PostUser(w http.ResponseWriter, r *http.Request) {
	uuid := ""
	decodedUser, err := server.DecodeJSON[User](r.Body)
	if err != nil {
		httperror.ErrorHandler(w, r, custom_error.UnknownError("user add", errors.New("incorrect request body")))
		return
	}
	domainUser := user.New(decodedUser.Name, decodedUser.Email, *decodedUser.Password)
	registerUser := command.NewRegisterUser(uuid, domainUser)
	if err := h.app.Command.RegisterUser.Handle(registerUser); err.Error() != "" {
		httperror.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusCreated)
}

func (h httpServer) PutUserUpdateEmail(w http.ResponseWriter, r *http.Request, params PutUserUpdateEmailParams) {
	uuid := ""
	updateUserEmail := command.NewUpdateUserEmail(uuid, params.Email)
	if err := h.app.Command.UpdateUserEmail.Handle(updateUserEmail); err.Error() != "" {
		httperror.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusCreated)
}

func (h httpServer) PutUserUpdateName(w http.ResponseWriter, r *http.Request, params PutUserUpdateNameParams) {
	uuid := ""
	updateUserName := command.NewUpdateUserName(uuid, params.Name)
	if err := h.app.Command.UpdateUserName.Handle(updateUserName); err.Error() != "" {
		httperror.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusCreated)
}

func (h httpServer) PutUserUpdatePassword(w http.ResponseWriter, r *http.Request, params PutUserUpdatePasswordParams) {
	uuid := ""
	updateUserPassword := command.NewUpdateUserPassword(uuid, params.Password)
	if err := h.app.Command.UpdateUserPassword.Handle(updateUserPassword); err.Error() != "" {
		httperror.ErrorHandler(w, r, err)
		return
	}
	server.Success(w, http.StatusCreated)
}
