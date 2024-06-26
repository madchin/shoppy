package adapters

import (
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"context"
	"database/sql"
)

type userDetailDto struct {
	uuid      string
	firstName string
	lastName  string
}

func NewUserDetailRepository(db *sql.DB) user.DetailRepository {
	return UserDetailRepository{db}
}

type UserDetailRepository struct {
	db *sql.DB
}

func (ur UserDetailRepository) Get(ctx context.Context, userUuid string) (user.UserDetail, custom_error.ContextError) {
	userDetailDto := userDetailDto{}
	err := ur.db.QueryRow("SELECT * FROM UserDetails WHERE uuid=?", userUuid).Scan(&userDetailDto.uuid, &userDetailDto.firstName, &userDetailDto.lastName)
	if err == sql.ErrNoRows {
		return user.UserDetail{}, custom_error.NewPersistenceError("user detail retrieve", "user with provided uuid not found")
	}
	domainUserDetail := userDetailDto.mapToDomainUserDetail()
	return domainUserDetail, custom_error.ContextError{}
}

func (ur UserDetailRepository) Create(ctx context.Context, userUuid string, ud user.UserDetail, validateFn func(user.UserDetail) []error) custom_error.ContextError {
	if u, _ := ur.Get(ctx, userUuid); u.Exists() {
		return custom_error.NewPersistenceError("user add", "user details already exists")
	}
	errs := validateFn(ud)
	if len(errs) > 0 {
		return custom_error.NewValidationErrors("user detail create", errs)
	}

	if _, err := ur.db.Exec("INSERT INTO UserDetails (uuid, firstName, lastName) VALUES (?, ?)", userUuid, ud.FirstName(), ud.LastName()); err != nil {
		return custom_error.UnknownPersistenceError("user detail create")
	}
	return custom_error.ContextError{}
}

func (ur UserDetailRepository) UpdateFirstName(ctx context.Context, userUuid string, name string, validateFn func(user.UserDetail) error) custom_error.ContextError {
	ud, err := ur.Get(ctx, userUuid)
	if err.Error() != "" {
		return err
	}

	if err := validateFn(ud); err != nil {
		return custom_error.NewValidationError("user detail update first name", err.Error())
	}
	if _, err := ur.db.Exec("UPDATE UserDetail SET firstName=? WHERE uuid=?", ud.FirstName(), userUuid); err != nil {
		return custom_error.UnknownPersistenceError("user detail update first name")
	}
	return custom_error.ContextError{}
}

func (ur UserDetailRepository) UpdateLastName(ctx context.Context, userUuid string, name string, validateFn func(user.UserDetail) error) custom_error.ContextError {
	ud, err := ur.Get(ctx, userUuid)
	if err.Error() != "" {
		return err
	}
	if err := validateFn(ud); err != nil {
		return custom_error.NewValidationError("user detail update last name", err.Error())
	}
	if _, err := ur.db.Exec("UPDATE UserDetail SET lastName=? WHERE uuid=?", ud.LastName(), userUuid); err != nil {
		return custom_error.UnknownPersistenceError("user detail update last name")
	}
	return custom_error.ContextError{}
}

func (ur UserDetailRepository) Delete(ctx context.Context, userUuid string) custom_error.ContextError {
	_, err := ur.Get(ctx, userUuid)
	if err.Error() != "" {
		return err
	}
	if _, err := ur.db.Exec("DELETE FROM UserDetail WHERE uuid=?", userUuid); err != nil {
		return custom_error.UnknownPersistenceError("user detail deletion")
	}
	return custom_error.ContextError{}
}

func (ud userDetailDto) mapToDomainUserDetail() user.UserDetail {
	return user.NewUserDetail(ud.firstName, ud.lastName)
}
