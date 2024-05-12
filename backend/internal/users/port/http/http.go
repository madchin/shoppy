package ports

import (
	"backend/internal/users/app"
	"backend/internal/users/port/http/handler"
)

type HttpServer struct {
	User handler.User
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{
		User: handler.NewUser(app),
	}
}
