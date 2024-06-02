package server

import (
	custom_error "backend/internal/common/errors"
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

func NewHttpError(context string, desc string) HttpError {
	return HttpError{context, desc}
}

func ParseCustomErrToHttpErrors(customError custom_error.ContextError) []HttpError {
	httpErrs := make([]HttpError, 0)
	for _, cerr := range customError.Errors() {
		httpErrs = append(httpErrs, NewHttpError(customError.Type().String(), cerr.Error()))
	}
	return httpErrs
}

func ErrorResponse(w http.ResponseWriter, status int, err ...HttpError) {
	w.Header().Set("Content-Type", "application/json")
	msg, _ := json.Marshal(err)
	w.WriteHeader(status)
	w.Write(msg)
}

func NotFound(w http.ResponseWriter, err ...HttpError) {
	ErrorResponse(w, http.StatusNotFound, err...)
}

func Unauthorized(w http.ResponseWriter, err ...HttpError) {
	ErrorResponse(w, http.StatusUnauthorized, err...)
}

func BadRequest(w http.ResponseWriter, err ...HttpError) {
	ErrorResponse(w, http.StatusBadRequest, err...)
}

func Internal(w http.ResponseWriter, err ...HttpError) {
	ErrorResponse(w, http.StatusInternalServerError, err...)
}
