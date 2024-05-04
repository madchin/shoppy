package handler

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Message string
	Code    int
}

var GenericError Error = Error{Code: http.StatusInternalServerError, Message: "Oops! something went wrong"}

func ErrorMsg(w http.ResponseWriter, status int, msg string) {
	parsedMsg, err := json.Marshal(Error{Code: status, Message: msg})
	if err != nil {
		msg, _ := json.Marshal(GenericError)
		w.WriteHeader(GenericError.Code)
		w.Write(msg)
		return
	}
	w.WriteHeader(status)
	w.Write(parsedMsg)
}
