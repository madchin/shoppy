package user

import (
	"errors"

	"github.com/hashicorp/go-multierror"
)

type UserDetail struct {
	firstName string
	lastName  string
}

var errFirstNameEmpty = errors.New("User first name is empty")
var errLastNameEmpty = errors.New("User last name is empty")

func (u UserDetail) Validate() error {
	err := u.validateFirstName()
	err = multierror.Append(err, u.validateLastName())
	return err.(*multierror.Error).ErrorOrNil()
}

func (u UserDetail) IsProvided() bool {
	return u.Validate() != nil
}

func (u UserDetail) validateFirstName() (err error) {
	if u.firstName == "" {
		err = errFirstNameEmpty
	}
	return err
}

func (u UserDetail) validateLastName() (err error) {
	if u.lastName == "" {
		err = errLastNameEmpty
	}
	return err
}
