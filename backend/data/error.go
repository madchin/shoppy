package data

import (
	"fmt"
)

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

func (e ErrEmptyEmail) Error() string {
	return "User email is empty"
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
