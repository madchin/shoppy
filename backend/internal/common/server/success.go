package server

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func SuccessWithBody(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	msg, _ := json.Marshal(body)
	w.WriteHeader(status)
	w.Write(msg)
}
