package user

import (
	"errors"
	"regexp"
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

func New(name string, email string) User {
	return User{name: name, email: email}
}

func (u User) Name() string {
	return u.name
}

func (u User) Email() string {
	return u.email
}

func (u User) Validate() (errs []error) {
	errs = u.ValidateName()
	errs = append(errs, u.ValidateEmail()...)
	return
}

func (u User) ValidateName() (errs []error) {
	if u.name == "" {
		errs = append(errs, errors.New("empty-name"))
	}
	if len(u.name) < nameMinLength {
		errs = append(errs, errors.New("short-name"))
	}
	if len(u.name) > nameMaxLength {
		errs = append(errs, errors.New("long-name"))
	}
	return
}

func (u User) ValidateEmail() (errs []error) {
	if u.email == "" {
		errs = append(errs, errors.New("empty-email"))
	}
	if len(u.email) > emailMaxLength {
		errs = append(errs, errors.New("long-email"))
	}
	ok, _ := regexp.MatchString(emailRegex, u.email)
	if !ok {
		errs = append(errs, errors.New("match-email"))
	}
	return
}
