package adapters

import (
	common_adapter "backend/internal/common/adapters"
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"database/sql"
)

type UserDetailDTO struct {
	uuid      string
	firstName string
	lastName  string
}

type UserDetailRepository struct {
	db *sql.DB
}

func (ur UserDetailRepository) Get(userUuid string) (user.UserDetail, custom_error.ContextError) {
	userDetail := UserDetailDTO{}
	err := ur.db.QueryRow("SELECT * FROM UserDetails WHERE uuid=?", userUuid).Scan(&userDetail.firstName, &userDetail.lastName)
	if err == sql.ErrNoRows {
		return user.UserDetail{}, custom_error.NewPersistenceError("user detail retrieve", "user with provided uuid not found")
	}
	domainUserDetail := user.NewUserDetail(userDetail.firstName, userDetail.lastName)
	return domainUserDetail, custom_error.ContextError{}
}

func (ur UserDetailRepository) Create(userUuid string, ud user.UserDetail, createFn func(user.UserDetail) (user.UserDetail, []error)) (user.UserDetail, custom_error.ContextError) {
	ud, errs := createFn(ud)
	if len(errs) > 0 {
		return user.UserDetail{}, custom_error.NewValidationErrors("user detail create", errs)
	}

	if _, err := ur.db.Exec("INSERT INTO UserDetails (uuid, firstName, lastName) VALUES (?, ?)", userUuid, ud.FirstName(), ud.LastName()); err != nil {
		if common_adapter.IsDuplicateEntryError(err) {
			return user.UserDetail{}, custom_error.NewPersistenceError("user detail create", "user with provided uuid already exists")
		}
		return user.UserDetail{}, custom_error.UnknownPersistenceError("user detail create")
	}
	return ud, custom_error.ContextError{}
}

func (ur UserDetailRepository) UpdateFirstName(userUuid string, name string, updateFn func(user.UserDetail) (user.UserDetail, []error)) custom_error.ContextError {
	ud, err := ur.Get(userUuid)
	if err.Error() != "" {
		return err
	}
	ud, errs := updateFn(ud)
	if len(errs) > 0 {
		return custom_error.NewValidationErrors("user detail update first name", errs)
	}
	if _, err := ur.db.Exec("UPDATE UserDetail SET firstName=? WHERE uuid=?", ud.FirstName(), userUuid); err != nil {
		return custom_error.UnknownPersistenceError("user detail update first name")
	}
	return custom_error.ContextError{}
}

func (ur UserDetailRepository) UpdateLastName(userUuid string, name string, updateFn func(user.UserDetail) (user.UserDetail, []error)) custom_error.ContextError {
	ud, err := ur.Get(userUuid)
	if err.Error() != "" {
		return err
	}
	ud, errs := updateFn(ud)
	if len(errs) > 0 {
		return custom_error.NewValidationErrors("user detail update last name", errs)
	}
	if _, err := ur.db.Exec("UPDATE UserDetail SET lastName=? WHERE uuid=?", ud.LastName(), userUuid); err != nil {
		return custom_error.UnknownPersistenceError("user detail update last name")
	}
	return custom_error.ContextError{}
}

func (ur UserDetailRepository) Delete(userUuid string) custom_error.ContextError {
	_, err := ur.Get(userUuid)
	if err.Error() != "" {
		return err
	}
	if _, err := ur.db.Exec("DELETE FROM UserDetail WHERE uuid=?", userUuid); err != nil {
		return custom_error.UnknownPersistenceError("user detail deletion")
	}
	return custom_error.ContextError{}
}
