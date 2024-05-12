package handler

import (
	"backend/internal/common/server"
	"backend/internal/users/app"
	"backend/internal/users/app/query"
	"net/http"
)

type User struct {
	app app.Application
}

func NewUser(app app.Application) User {
	return User{app: app}
}

func (u User) Get(w http.ResponseWriter, r *http.Request) {
	queryUser, err := u.app.Query.RetrieveUser.Handle(query.RetrieveUser{Uuid: "elo"})
	if err != nil {
		server.ErrorResponse(w, server.HttpErrInternal)
		return
	}
	server.SuccessResponse(w, queryUser, http.StatusOK)
}
