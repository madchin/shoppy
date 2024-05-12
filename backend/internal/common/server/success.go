package server

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse[T any](w http.ResponseWriter, body T, status int) {
	w.Header().Set("Content-Type", "application/json")
	msg, err := json.Marshal(body)
	if err != nil {
		errorJsonResponse(w)
		return
	}
	w.WriteHeader(status)
	w.Write(msg)
}
