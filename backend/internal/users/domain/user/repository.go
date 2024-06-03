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
	Get(userUuid string) (UserDetail, custom_error.ContextError)
	Create(
		userUuid string,
		userDetail UserDetail,
		createFn func(UserDetail) (UserDetail, []error),
	) custom_error.ContextError
	UpdateFirstName(
		userUuid string,
		name string,
		updateFn func(UserDetail) (UserDetail, error),
	) custom_error.ContextError
	UpdateLastName(
		userUuid string,
		name string,
		updateFn func(UserDetail) (UserDetail, error),
	) custom_error.ContextError
	Delete(
		userUuid string,
	) custom_error.ContextError
}

type PhoneRepository interface {
	UpdateNumber(
		userUuid string,
		number int,
		updateFn func(Phone) (Phone, error),
	) error
}
