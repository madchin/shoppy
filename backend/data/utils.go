package data

type validationError struct {
	EmptyEmail *emptyEmail
}

type emptyEmail struct{}

var ValidationError = &validationError{
	EmptyEmail: &emptyEmail{},
}

func (e *emptyEmail) Error() string {
	return "User email is empty"
}
