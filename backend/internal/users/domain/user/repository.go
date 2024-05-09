package user

type Repository interface {
	Create(
		uuid string,
		user User,
		createFn func(User) (User, error),
	) error
	UpdateName(
		uuid string,
		user User,
		updateFn func(User) (User, error),
	)
	UpdateEmail(
		uuid string,
		user User,
		updateFn func(User) (User, error),
	)
	Delete(
		uuid string,
		user User,
		deleteFn func(User) error,
	)
}

type DetailRepository interface {
	UpdateFirstName(
		userUuid string,
		name string,
		updateFn func(UserDetail) (UserDetail, error),
	)
	UpdateLastName(
		userUuid string,
		name string,
		updateFn func(UserDetail) (UserDetail, error),
	)
}

type PhoneRepository interface {
	UpdateNumber(
		userUuid string,
		number int,
		updateFn func(Phone) (Phone, error),
	)
}
