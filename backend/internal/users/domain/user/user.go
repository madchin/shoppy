package user

import (
	"errors"
	"regexp"
)

type User struct {
	uuid     string
	name     string
	email    string
	password string
}

const (
	nameMinLength     = 6
	nameMaxLength     = 32
	passwordMinLength = 12
	emailRegex        = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
	emailMaxLength    = 320
)

func NewUser(uuid string, password string, name string, email string) User {
	return User{uuid: uuid, password: password, name: name, email: email}
}

func (u User) Name() string {
	return u.name
}

func (u User) Email() string {
	return u.email
}

func (u User) Uuid() string {
	return u.uuid
}

func (u User) Password() string {
	return u.password
}

func (u User) Validate() (errs []error) {
	errs = u.ValidateName()
	errs = append(errs, u.ValidateEmail()...)
	errs = append(errs, u.ValidatePassword()...)
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
func (u User) ValidatePassword() (errs []error) {
	if u.password == "" {
		errs = append(errs, errors.New("empty-password"))
	}
	if len(u.password) < passwordMinLength {
		errs = append(errs, errors.New("short-password"))
	}
	return
}

func (u User) IsPasswordEqual(password string) (err error) {
	if u.password == password {
		return nil
	}
	return errors.New("password is not correct")
}
