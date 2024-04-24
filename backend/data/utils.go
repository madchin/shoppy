package data

import "fmt"

type emptyEmail struct{}

type missingEnv struct {
	Keys []string
}

type missingUuid struct{}

type missingFirstName struct{}

type missingLastName struct{}

type errDbConfig struct {
	MissingEnv *missingEnv
}

type errUser struct {
	User       *User
	EmptyEmail *emptyEmail
}

type errUserDetail struct {
	MissingUuid      *missingUuid
	UserDetail       *UserDetail
	MissingFirstName *missingFirstName
	MissingLastName  *missingLastName
}

var ErrDbConfig = &errDbConfig{}

var ErrUser = &errUser{}

var ErrUserDetail = &errUserDetail{}

func (e *missingEnv) Add(env string) {
	e.Keys = append(e.Keys, env)
}

func (e *emptyEmail) Error() string {
	return "User email is empty"
}

func (e *missingEnv) Error() string {
	var missCount int
	var missingEnvs string
	for _, env := range e.Keys {
		missCount++
		missingEnvs += fmt.Sprintf("%s, ", env)
	}
	return fmt.Sprintf("%d count of envs are missing: %s", missCount, missingEnvs)
}

func (e *missingUuid) Error() string {
	return fmt.Sprintf("Missing uuid for user with firstName: %s and lastName: %s", ErrUserDetail.UserDetail.FirstName, ErrUserDetail.UserDetail.LastName)
}

func (e *missingFirstName) Error() string {
	return fmt.Sprintf("Missing first name for user with uuid: %s and lastName: %s", ErrUserDetail.UserDetail.Uuid, ErrUserDetail.UserDetail.LastName)
}

func (e *missingLastName) Error() string {
	return fmt.Sprintf("Missing last name for user with uuid: %s and lastName: %s", ErrUserDetail.UserDetail.Uuid, ErrUserDetail.UserDetail.LastName)
}
