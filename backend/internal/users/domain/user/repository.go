package user

type Repository interface {
	Create(
		uuid string,
		user User,
		createFn func(User) (User, error),
	) error
	Get(
		uuid string,
	) (User, error)
	UpdateName(
		uuid string,
		user User,
		updateFn func(User) (User, error),
	) error
	UpdateEmail(
		uuid string,
		user User,
		updateFn func(User) (User, error),
	) error
	Delete(
		uuid string,
		deleteFn func(User) error,
	) error
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
