package ports

import (
	"backend/internal/users/app"
	"backend/internal/users/app/command"
	"backend/internal/users/port/http/httperror"
	"net/http"
)

type httpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) ServerInterface {
	return httpServer{app}
}

func (h httpServer) DeleteUser(w http.ResponseWriter, r *http.Request) {
	uuid := "r."
	user := command.NewDeleteUser(uuid)
	err := h.app.Command.DeleteUser.Handle(user)
	if err.Error() != "" {
		httperror.ErrorHandler(w, r, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h httpServer) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (h httpServer) PostUser(w http.ResponseWriter, r *http.Request) {

}

func (h httpServer) PutUserUpdateEmail(w http.ResponseWriter, r *http.Request, params PutUserUpdateEmailParams) {

}

func (h httpServer) PutUserUpdateName(w http.ResponseWriter, r *http.Request, params PutUserUpdateNameParams) {

}

func (h httpServer) PutUserUpdatePassword(w http.ResponseWriter, r *http.Request, params PutUserUpdatePasswordParams) {

}
