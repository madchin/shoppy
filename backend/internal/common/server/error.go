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

func ErrorHandler(w http.ResponseWriter, r *http.Request, err custom_error.ContextError) {
	httpErrs := parseCustomErrToHttpErrors(err)
	switch err.Type() {
	case custom_error.ErrorTypeAuthorization:
		unauthorized(w, httpErrs...)
	case custom_error.ErrorTypeValidation:
		badRequest(w, httpErrs...)
	case custom_error.ErrorTypePersistence:
		badRequest(w, httpErrs...)
	default:
		internal(w, httpErrs...)
	}
}

func parseCustomErrToHttpErrors(customError custom_error.ContextError) []HttpError {
	httpErrs := make([]HttpError, 0)
	for _, cerr := range customError.Errors() {
		httpErrs = append(httpErrs, NewHttpError(customError.Type().String(), cerr.Error()))
	}
	return httpErrs
}

func errorResponse(w http.ResponseWriter, status int, err ...HttpError) {
	w.Header().Set("Content-Type", "application/json")
	msg, _ := json.Marshal(err)
	w.WriteHeader(status)
	w.Write(msg)
}

func unauthorized(w http.ResponseWriter, err ...HttpError) {
	errorResponse(w, http.StatusUnauthorized, err...)
}

func badRequest(w http.ResponseWriter, err ...HttpError) {
	errorResponse(w, http.StatusBadRequest, err...)
}

func internal(w http.ResponseWriter, err ...HttpError) {
	errorResponse(w, http.StatusInternalServerError, err...)
}
