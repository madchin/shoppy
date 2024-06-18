package user

import (
	custom_error "backend/internal/common/errors"
	"errors"
)

type UserDetail struct {
	firstName string
	lastName  string
}

func (ud UserDetail) FirstName() string {
	return ud.firstName
}

func (ud UserDetail) LastName() string {
	return ud.lastName
}

func NewUserDetail(firstName string, lastName string) UserDetail {
	return UserDetail{firstName, lastName}
}
func (u UserDetail) Validate() (errs []error) {
	errs = custom_error.AppendError(errs, u.ValidateFirstName())
	errs = custom_error.AppendError(errs, u.ValidateLastName())
	return
}

func (u UserDetail) Exists() bool {
	return u.firstName != "" || u.lastName != ""
}

func (u UserDetail) ValidateFirstName() (err error) {
	if u.firstName == "" {
		err = errors.New("User first name is empty")
	}
	return
}

func (u UserDetail) ValidateLastName() (err error) {
	if u.lastName == "" {
		err = errors.New("User last name is empty")
	}
	return
}
