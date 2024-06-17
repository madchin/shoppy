package custom_error

// I wanna have typed errors returned from domain layer in order to recognize them, i can return simply error type from domain because i will resolve all types in persistence
// return them in persistence layer as map[ERrortype][]errors
// in app layer ik would like to take this map and parse it to another type, to add an context to an operation (where it happened)
// then i can consume it in port layer via switch case and return revelant response to user (array json with all errors. it can be multiple validation errors or one persistence error)
// in case of one json object, in case of > 1 array with objects
import (
	"errors"
	"fmt"
	"strings"
)

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown       = ErrorType{"unknown"}
	ErrorTypePersistence   = ErrorType{"persistence"}
	ErrorTypeValidation    = ErrorType{"validation"}
	ErrorTypeAuthorization = ErrorType{"authorization"}
	ErrorTypeConfiguration = ErrorType{"config"}
)

var (
	errUnknown = errors.New("unknown")
)

type ContextError struct {
	context   string
	errorType ErrorType
	errors    []error
}

func (e ErrorType) String() string {
	return e.t
}

func (e ContextError) Type() ErrorType {
	return e.errorType
}

func (e ContextError) Errors() []error {
	return e.errors
}

func (e ContextError) Context() string {
	return e.context
}

func (e ContextError) Error() string {
	var errMsg []string
	for _, err := range e.errors {
		errMsg = append(errMsg, err.Error())
	}
	if len(errMsg) > 0 {
		return fmt.Sprintf("In %s operation error: %s  error message: %s", e.context, e.errorType.t, strings.Join(errMsg, ", "))
	}
	return ""
}

func newContextError(context string, errorType ErrorType, errors []error) ContextError {
	return ContextError{context, errorType, errors}
}

func NewPersistenceError(context string, message string) ContextError {
	return newContextError(context, ErrorTypePersistence, []error{errors.New(message)})
}

func NewValidationErrors(context string, errs []error) ContextError {
	return newContextError(context, ErrorTypeValidation, errs)
}

func NewValidationError(context string, message string) ContextError {
	return newContextError(context, ErrorTypeValidation, []error{errors.New(message)})
}

func NewAuthorizationError(context string, message string) ContextError {
	return newContextError(context, ErrorTypeAuthorization, []error{errors.New(message)})
}

func NewConfigurationError(context string, message string) ContextError {
	return newContextError(context, ErrorTypeConfiguration, []error{errors.New(message)})
}

func UnknownError(context string) ContextError {
	return newContextError(context, ErrorTypeUnknown, []error{errUnknown})
}

func UnknownPersistenceError(context string) ContextError {
	return newContextError(context, ErrorTypePersistence, []error{errUnknown})
}
