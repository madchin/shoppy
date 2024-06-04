package user

import custom_error "backend/internal/common/errors"

type Repository interface {
	Create(
		uuid string,
		user User,
		validateFn func(User) []error,
	) custom_error.ContextError
	Get(uuid string) (User, custom_error.ContextError)
	UpdateName(
		uuid string,
		name string,
		validateFn func(User) []error,
	) custom_error.ContextError
	UpdateEmail(
		uuid string,
		email string,
		validateFn func(User) []error,
	) custom_error.ContextError
	UpdatePassword(
		uuid string,
		password string,
		validateFn func(User) []error,
	) custom_error.ContextError
	Delete(uuid string) custom_error.ContextError
}

type DetailRepository interface {
	Get(userUuid string) (UserDetail, custom_error.ContextError)
	Create(
		userUuid string,
		userDetail UserDetail,
		validateFn func(UserDetail) []error,
	) custom_error.ContextError
	UpdateFirstName(
		userUuid string,
		name string,
		validateFn func(UserDetail) error,
	) custom_error.ContextError
	UpdateLastName(
		userUuid string,
		name string,
		validateFn func(UserDetail) error,
	) custom_error.ContextError
	Delete(
		userUuid string,
	) custom_error.ContextError
}

type PhoneRepository interface {
	Create(
		userUuid string,
		phone Phone,
		validateFn func(Phone) []error,
	) custom_error.ContextError
	Get(
		userUuid string,
	) (Phones, custom_error.ContextError)
	Update(
		userUuid string,
		prevNumber string,
		phone Phone,
		validateFn func(Phone) []error,
	) custom_error.ContextError
	DeletePhone(
		userUuid string,
		phone Phone,
		validateFn func(Phones) error,
	) custom_error.ContextError
	DeleteAll(
		userUuid string,
	) custom_error.ContextError
}
