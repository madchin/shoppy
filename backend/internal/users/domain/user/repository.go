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
	Create(
		userUuid string,
		number int,
		createFn func(Phone) (Phone, []error),
	) custom_error.ContextError
	Get(
		userUuid string,
	) (Phones, custom_error.ContextError)
	Update(
		userUuid string,
		prevNumber string,
		phone Phone,
		updateFn func(Phone) (Phone, []error),
	) custom_error.ContextError
	DeletePhone(
		userUuid string,
		phone Phone,
		deleteFn func(Phones) (Phones, error),
	) custom_error.ContextError
	DeleteAll(
		userUuid string,
	) custom_error.ContextError
}
