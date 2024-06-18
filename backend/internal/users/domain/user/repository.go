package user

import (
	custom_error "backend/internal/common/errors"
	"context"
)

type Repository interface {
	Create(
		ctx context.Context,
		uuid string,
		user User,
		validateFn func(User) []error,
	) custom_error.ContextError
	GetByUuid(ctx context.Context, uuid string) (User, custom_error.ContextError)
	GetByEmail(ctx context.Context, email string, validatePasswordFn func(User) error) (User, custom_error.ContextError)
	FindByEmail(ctx context.Context, email string) (User, custom_error.ContextError)
	UpdateName(
		ctx context.Context,
		uuid string,
		name string,
		validateFn func(User) []error,
	) custom_error.ContextError
	UpdateEmail(
		ctx context.Context,
		uuid string,
		email string,
		validateFn func(User) []error,
	) custom_error.ContextError
	UpdatePassword(
		ctx context.Context,
		uuid string,
		password string,
		validateFn func(User) []error,
	) custom_error.ContextError
	Delete(ctx context.Context, uuid string) custom_error.ContextError
}

type DetailRepository interface {
	Get(ctx context.Context, userUuid string) (UserDetail, custom_error.ContextError)
	Create(
		ctx context.Context,
		userUuid string,
		userDetail UserDetail,
		validateFn func(UserDetail) []error,
	) custom_error.ContextError
	UpdateFirstName(
		ctx context.Context,
		userUuid string,
		name string,
		validateFn func(UserDetail) error,
	) custom_error.ContextError
	UpdateLastName(
		ctx context.Context,
		userUuid string,
		name string,
		validateFn func(UserDetail) error,
	) custom_error.ContextError
	Delete(
		ctx context.Context,
		userUuid string,
	) custom_error.ContextError
}

type PhoneRepository interface {
	Create(
		ctx context.Context,
		userUuid string,
		phone Phone,
		validateFn func(Phone) []error,
	) custom_error.ContextError
	Get(
		ctx context.Context,
		userUuid string,
	) (Phones, custom_error.ContextError)
	Update(
		ctx context.Context,
		userUuid string,
		prevNumber string,
		phone Phone,
		validateFn func(Phone) []error,
	) custom_error.ContextError
	DeletePhone(
		ctx context.Context,
		userUuid string,
		phone Phone,
		validateFn func(Phones) error,
	) custom_error.ContextError
	DeleteAll(
		ctx context.Context,
		userUuid string,
	) custom_error.ContextError
}

type AddressRepository interface {
	Create(
		ctx context.Context,
		userUuid string,
		address Address,
		validateFn func(Address) []error,
	) custom_error.ContextError
	Get(
		ctx context.Context,
		userUuid string,
	) (Addresses, custom_error.ContextError)
	Update(
		ctx context.Context,
		userUuid string,
		addressStreet string,
		address Address,
		validateFn func(Address) []error,
	) custom_error.ContextError
	DeleteAddress(
		ctx context.Context,
		userUuid string,
		street string,
	) custom_error.ContextError
	DeleteAll(
		ctx context.Context,
		userUuid string,
	) custom_error.ContextError
}
