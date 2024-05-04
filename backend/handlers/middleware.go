package handler

import (
	"net/http"
)

func SessionMiddleware(userService userService, handler UserHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check for session
		handler(userService, "", w, r)
	}
}
