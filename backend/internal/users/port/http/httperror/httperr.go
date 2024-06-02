package httperror

import (
	custom_error "backend/internal/common/errors"
	"backend/internal/common/server"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err custom_error.ContextError) {
	httpErrs := server.ParseCustomErrToHttpErrors(err)
	switch err.Type() {
	case custom_error.ErrorTypeAuthorization:
		server.Unauthorized(w, httpErrs...)
	case custom_error.ErrorTypeValidation:
	case custom_error.ErrorTypePersistence:
		server.BadRequest(w, httpErrs...)
	default:
		server.Internal(w, httpErrs...)
	}
}
