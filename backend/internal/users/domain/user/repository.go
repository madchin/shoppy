package user

import custom_error "backend/internal/common/errors"

type Repository interface {
	Create(
		uuid string,
		user User,
		createFn func(User) (User, []error),
	) custom_error.ContextError
	Get(uuid string) (User, custom_error.ContextError)
	UpdateName(
		uuid string,
		name string,
		updateFn func(User) (User, []error),
	) custom_error.ContextError
	UpdateEmail(
		uuid string,
		email string,
		updateFn func(User) (User, []error),
	) custom_error.ContextError
	UpdatePassword(
		uuid string,
		password string,
		updateFn func(User) (User, []error),
	) custom_error.ContextError
	Delete(uuid string) custom_error.ContextError
}

type DetailRepository interface {
	UpdateFirstName(
		userUuid string,
		name string,
		updateFn func(UserDetail) (UserDetail, error),
	) error
	UpdateLastName(
		userUuid string,
		name string,
		updateFn func(UserDetail) (UserDetail, error),
	) error
}

type PhoneRepository interface {
	UpdateNumber(
		userUuid string,
		number int,
		updateFn func(Phone) (Phone, error),
	) error
}
