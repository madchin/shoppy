package handler

import (
	"backend/data"
	"net/http"
)

func SessionMiddleware(userRepo *data.UserRepository, handler UserHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check for session
		handler(userRepo, "", w, r)
	}
}
