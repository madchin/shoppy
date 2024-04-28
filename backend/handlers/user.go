package handler

import (
	"backend/data"
	"database/sql"
	"encoding/json"
	"net/http"
)

func get(db *sql.DB, uuid string, w http.ResponseWriter) {
	if uuid == "" {
		ErrorMsg(w, http.StatusUnauthorized, "Unauthorized to perform this action")
		return
	}
	user, err := data.GetUser(db, uuid)
	if err != nil {
		ErrorMsg(w, http.StatusBadRequest, "Unable to retrieve user")
		return
	}
	msg, err := json.Marshal(user)
	if err != nil {
		ErrorMsg(w, http.StatusInternalServerError, GenericError.Message)
		return
	}
	w.Write(msg)
}

func create(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	user := &data.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		ErrorMsg(w, http.StatusBadRequest, GenericError.Message)
		return
	}
	if err := user.Create(db); err != nil {
		ErrorMsg(w, http.StatusBadRequest, "Error occured during creating user")
		return
	}
	msg, err := json.Marshal(user)
	if err != nil {
		ErrorMsg(w, http.StatusInternalServerError, GenericError.Message)
		return
	}
	w.Write(msg)
}

func update(db *sql.DB, uuid string, w http.ResponseWriter, r *http.Request) {
	if uuid == "" {
		ErrorMsg(w, http.StatusUnauthorized, "Unauthorized to perform this action")
		return
	}
	incomingUser := &data.User{}
	if err := json.NewDecoder(r.Body).Decode(incomingUser); err != nil {
		ErrorMsg(w, http.StatusBadRequest, GenericError.Message)
		return
	}
	user, err := data.GetUser(db, uuid)
	if err != nil {
		ErrorMsg(w, http.StatusBadRequest, "User with provided id not exists")
		return
	}
	if user.Name != incomingUser.Name {
		if err = user.UpdateName(db); err != nil {
			ErrorMsg(w, http.StatusBadRequest, "Updating name failed")
			return
		}
	}
	if user.Email != incomingUser.Email {
		if err = user.UpdateEmail(db); err != nil {
			ErrorMsg(w, http.StatusBadRequest, "Updating email failed")
			return
		}
	}
	msg, err := json.Marshal(user)
	if err != nil {
		ErrorMsg(w, http.StatusInternalServerError, GenericError.Message)
		return
	}
	w.Write(msg)
}

func User(db *sql.DB, uuid string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		create(db, w, r)
	case "GET":
		get(db, uuid, w)
	case "PUT":
		update(db, uuid, w, r)
	default:
		ErrorMsg(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
