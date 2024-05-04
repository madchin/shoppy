package data

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

type ErrMissingEnv struct {
	Keys []string
}

type ErrMissingUuid struct{}

type ErrMissingPhoneNumber struct {
	Phone
}

type ErrMissingFirstName struct {
	userDetail UserDetail
}

type ErrMissingLastName struct {
	userDetail UserDetail
}

type ErrEmptyEmail struct {
	user User
}

func (e *ErrMissingEnv) Add(env string) {
	e.Keys = append(e.Keys, env)
}

func (e ErrEmptyEmail) Error() string {
	return "User email is empty"
}

func (e ErrMissingEnv) Error() string {
	var missCount int
	var missingEnvs string
	for _, env := range e.Keys {
		missCount++
		missingEnvs += fmt.Sprintf("%s, ", env)
	}
	return fmt.Sprintf("%d count of envs are missing: %s", missCount, missingEnvs)
}

func (e ErrMissingUuid) Error() string {
	return "Missing uuid"
}

func (e ErrMissingFirstName) Error() string {
	return fmt.Sprintf("Missing first name for user with uuid: %s and lastName: %s", e.userDetail.Uuid, e.userDetail.FirstName)
}

func (e ErrMissingLastName) Error() string {
	return fmt.Sprintf("Missing last name for user with uuid: %s and lastName: %s", e.userDetail.Uuid, e.userDetail.LastName)
}

func (e ErrMissingPhoneNumber) Error() string {
	return fmt.Sprintf("Missing phone number for user with uuid %s", e.Phone.Uuid)
}

var UserColumns = []*sqlmock.Column{
	sqlmock.NewColumn("uuid").OfType("varchar(36)", uuid.NewString()).Nullable(false),
	sqlmock.NewColumn("name").OfType("varchar(255)", "randomName"),
	sqlmock.NewColumn("email").OfType("varchar(255)", "email@email.com").Nullable(false),
}
