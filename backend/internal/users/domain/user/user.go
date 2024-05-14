package user

import (
	"errors"
	"regexp"

	"backend/internal/common/domain"

	"github.com/hashicorp/go-multierror"
)

type User struct {
	name  string
	email string
}

const (
	nameMinLength  = 6
	nameMaxLength  = 32
	emailRegex     = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	emailMaxLength = 320
)

var errNameEmpty = errors.New("User name is empty")
var errNameMinLength = errors.New("User name is too short")
var errNameMaxLength = errors.New("User name is too long")
var errEmailEmpty = errors.New("User email is empty")
var errEmailMaxLength = errors.New("User email is too long")
var errEmailNotMatch = errors.New("User email have wrong format")

func New(name string, email string) User {
	return User{name: name, email: email}
}

func (u User) Validate() error {
	var err error
	err = u.ValidateName()
	err = multierror.Append(err, u.ValidateEmail())
	return err.(*multierror.Error).ErrorOrNil()
}

func (u User) ValidateName() (err error) {
	if u.name == "" {
		err = multierror.Append(err, errNameEmpty)
	}
	if len(u.name) < nameMinLength {
		err = multierror.Append(err, errNameMinLength)
	}
	if len(u.name) > nameMaxLength {
		err = multierror.Append(err, errNameMaxLength)
	}
	return err
}

func (u User) ValidateEmail() (err error) {
	if u.email == "" {
		err = multierror.Append(err, errEmailEmpty)
	}
	if len(u.email) > emailMaxLength {
		err = multierror.Append(err, errEmailMaxLength)
	}
	ok, rerr := regexp.MatchString(emailRegex, u.email)
	if rerr != nil {
		err = multierror.Append(err, domain.ErrInternal)
	}
	if !ok {
		err = multierror.Append(err, errEmailNotMatch)
	}
	return err
}
