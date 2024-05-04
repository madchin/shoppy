package handler

import (
	"backend/data"
	"encoding/json"
	"net/http"
	"strings"
)

type userService interface {
	GetUser(uuid string) (data.User, error)
	Create(user data.User) error
	UpdateName(user data.User) error
	UpdateEmail(user data.User) error
}

func get(userService userService, uuid string, w http.ResponseWriter) {
	user, err := userService.GetUser(uuid)
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

func create(userService userService, w http.ResponseWriter, r *http.Request) {
	user := &data.User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		ErrorMsg(w, http.StatusBadRequest, GenericError.Message)
		return
	}
	if err := userService.Create(*user); err != nil {
		ErrorMsg(w, http.StatusBadRequest, "Error occured during creating user")
		return
	}
	msg, err := json.Marshal(user)
	//branch not handled in tests, any possibility to occur? need to mock?
	if err != nil {
		ErrorMsg(w, http.StatusInternalServerError, GenericError.Message)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(msg)
}

func update(userService userService, uuid string, w http.ResponseWriter, r *http.Request) {
	if uuid == "" {
		ErrorMsg(w, http.StatusUnauthorized, "Unauthorized to perform this action")
		return
	}
	incomingUser := &data.User{}
	if err := json.NewDecoder(r.Body).Decode(incomingUser); err != nil {
		ErrorMsg(w, http.StatusBadRequest, GenericError.Message)
		return
	}
	user, err := userService.GetUser(uuid)
	if err != nil {
		ErrorMsg(w, http.StatusBadRequest, "User with provided id not exists")
		return
	}

	errorChannel := make(chan string, 2)
	go func(chan<- string) {
		if user.Name != incomingUser.Name {
			if err = userService.UpdateName(user); err != nil {
				errorChannel <- "Updating name failed"
				return
			}
			errorChannel <- ""
		}
	}(errorChannel)
	go func(chan<- string) {
		if user.Email != incomingUser.Email {
			if err = userService.UpdateEmail(user); err != nil {
				errorChannel <- "Updating email failed"
				return
			}
			errorChannel <- ""
		}
	}(errorChannel)
	var errorMessages []string
	for i := 0; i < 2; i++ {
		msg := <-errorChannel
		if msg != "" {
			errorMessages = append(errorMessages, msg)
		}
	}
	if len(errorMessages) > 0 {
		ErrorMsg(w, http.StatusBadRequest, strings.Join(errorMessages, ", "))
		return
	}
	msg, err := json.Marshal(user)
	if err != nil {
		ErrorMsg(w, http.StatusInternalServerError, GenericError.Message)
		return
	}
	w.Write(msg)
}

func User(service userService, uuid string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "POST":
		create(service, w, r)
	case "GET":
		get(service, uuid, w)
	case "PUT":
		update(service, uuid, w, r)
	default:
		ErrorMsg(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}
