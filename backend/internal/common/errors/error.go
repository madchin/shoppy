package custom_error

// I wanna have typed errors returned from domain layer in order to recognize them, i can return simply error type from domain because i will resolve all types in persistence
// return them in persistence layer as map[ERrortype][]errors
// in app layer ik would like to take this map and parse it to another type, to add an context to an operation (where it happened)
// then i can consume it in port layer via switch case and return revelant response to user (array json with all errors. it can be multiple validation errors or one persistence error)
// in case of one json object, in case of > 1 array with objects
import (
	"fmt"
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
	errMsg := ""
	for _, err := range e.errors {
		errMsg += err.Error()
	}
	return fmt.Sprintf("In %s operation error %T occured %s", e.context, e.errorType, errMsg)
}

type ErrMissingEnv struct {
	ContextError
	Keys      []string
	missCount int
}

func NewContextError(context string, errorType ErrorType, errors []error) ContextError {
	return ContextError{context, errorType, errors}
}

// FIXME
func NewErrMissingEnv(contextErr ContextError) *ErrMissingEnv {
	return &ErrMissingEnv{ContextError{}, []string{}, 0}
}

func (e *ErrMissingEnv) Add(env string) {
	e.Keys = append(e.Keys, env)
}

func (e *ErrMissingEnv) Error() string {
	missingEnvs := make([]string, 0)
	for _, env := range e.Keys {
		e.missCount++
		missingEnvs = append(missingEnvs, env)
	}

	return fmt.Sprintf("%d count of envs are missing: %s", e.missCount, missingEnvs)
}
