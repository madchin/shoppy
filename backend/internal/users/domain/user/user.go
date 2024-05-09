package user

import (
	"errors"
	"regexp"

	cerr "backend/internal/common/error"

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

func (u User) Validate() error {
	var err error
	err = u.validateName()
	err = multierror.Append(err, u.validateEmail())
	return err.(*multierror.Error).ErrorOrNil()
}

func (u User) validateName() (err error) {
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

func (u User) validateEmail() (err error) {
	if u.email == "" {
		err = multierror.Append(err, errEmailEmpty)
	}
	if len(u.email) > emailMaxLength {
		err = multierror.Append(err, errEmailMaxLength)
	}
	ok, rerr := regexp.MatchString(emailRegex, u.email)
	if rerr != nil {
		err = multierror.Append(err, cerr.ErrInternal)
	}
	if !ok {
		err = multierror.Append(err, errEmailNotMatch)
	}
	return err
}
