package data

import "fmt"

type errMissingEnv struct {
	Keys []string
}

var ErrMissingEnv = &errMissingEnv{}

type ErrMissingUuid struct {
	user       *User
	userDetail *UserDetail
}

type ErrMissingPhoneNumber struct {
	phone *Phone
}

type ErrMissingFirstName struct {
	userDetail *UserDetail
}

type ErrMissingLastName struct {
	userDetail *UserDetail
}

type ErrEmptyEmail struct {
	user *User
}

func (e *errMissingEnv) Add(env string) {
	e.Keys = append(e.Keys, env)
}

func (e *ErrEmptyEmail) Error() string {
	return "User email is empty"
}

func (e *errMissingEnv) Error() string {
	var missCount int
	var missingEnvs string
	for _, env := range e.Keys {
		missCount++
		missingEnvs += fmt.Sprintf("%s, ", env)
	}
	return fmt.Sprintf("%d count of envs are missing: %s", missCount, missingEnvs)
}

func (e *ErrMissingUuid) Error() string {
	if e.userDetail != nil {
		return fmt.Sprintf("Missing uuid for user with firstName: %s and lastName: %s", e.userDetail.FirstName, e.userDetail.LastName)
	}
	if e.user != nil {
		return fmt.Sprintf("Missing uuid for user with Name: %s and Email: %s", e.user.Name, e.user.Email)
	}
	return "Missing uuid"
}

func (e *ErrMissingFirstName) Error() string {
	return fmt.Sprintf("Missing first name for user with uuid: %s and lastName: %s", e.userDetail.Uuid, e.userDetail.FirstName)
}

func (e *ErrMissingLastName) Error() string {
	return fmt.Sprintf("Missing last name for user with uuid: %s and lastName: %s", e.userDetail.Uuid, e.userDetail.LastName)
}

func (e *ErrMissingPhoneNumber) Error() string {
	return fmt.Sprintf("Missing phone number for user with uuid %s", e.phone.Uuid)
}
