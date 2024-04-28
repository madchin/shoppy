package handler

import (
	"database/sql"
	"net/http"
)

func SessionMiddleware(db *sql.DB, handler HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check for session
		uuid := ""
		handler(db, uuid, w, r)
	}
}
