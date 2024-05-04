package data

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

var userDetailColumns = []*sqlmock.Column{
	sqlmock.NewColumn("uuid").OfType("varchar(36)", uuid.NewString()).Nullable(false),
	sqlmock.NewColumn("firstName").OfType("varchar(255)", "firstName"),
	sqlmock.NewColumn("lastName").OfType("varchar(255)", "lastName").Nullable(false),
}

func TestGetUserDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	repo := NewUserDetailRepository(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("Should get user details", func(t *testing.T) {
		uuid := uuid.New().String()
		userDetail := UserDetail{Uuid: uuid, FirstName: "firstName", LastName: "lastName"}
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM UserDetails WHERE uuid=?")).WithArgs(uuid).WillReturnRows(sqlmock.NewRowsWithColumnDefinition(userDetailColumns...).AddRow(userDetail.Uuid, userDetail.FirstName, userDetail.LastName))
		selectedUserDetails, err := repo.GetUserDetails(uuid)
		if err != nil {
			t.Fatalf(fmt.Sprintf("%v", err))
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
		if selectedUserDetails.FirstName != userDetail.FirstName {
			t.Fatalf("selected user detail first name is not equal user detail inserted in db, expected: %s, actual: %s: ", userDetail.FirstName, selectedUserDetails.FirstName)
		}
		if selectedUserDetails.Uuid != userDetail.Uuid {
			t.Fatalf("selected user detail uuid is not equal user detail inserted in db, expected: %s, actual: %s: ", userDetail.Uuid, selectedUserDetails.Uuid)
		}
		if selectedUserDetails.LastName != userDetail.LastName {
			t.Fatalf("selected user detail last name is not equal user detail inserted in db, expected: %s, actual: %s: ", userDetail.LastName, selectedUserDetails.LastName)
		}
	})

	t.Run("Should NOT get user details when user uuid is empty", func(t *testing.T) {
		_, err := repo.GetUserDetails("")
		if err != err.(ErrMissingUuid) {
			t.Fatalf(fmt.Sprintf("An error different than expected occured, actual error: %v", err))
		}
	})
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	repo := NewUserDetailRepository(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("Should create user details", func(t *testing.T) {
		uuid := uuid.New().String()
		userDetail := UserDetail{Uuid: uuid, FirstName: "firstName", LastName: "lastName"}
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO UserDetails (uuid, firstName, lastName) VALUES (?, ?, ?)")).WithArgs(userDetail.Uuid, userDetail.FirstName, userDetail.LastName).WillReturnResult(sqlmock.NewResult(0, 1))
		err = repo.Create(userDetail)

		if err != nil {
			t.Fatalf("User Details record has not been created, actual error: %v", err)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})

	t.Run("Should NOT create user details when uuid is not provided", func(t *testing.T) {
		userDetail := UserDetail{FirstName: "firstName", LastName: "lastName"}
		err = repo.Create(userDetail)
		if err != err.(ErrMissingUuid) {
			t.Fatalf("An error different than expected occured, actual error: %v", err)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})

}

func TestUpdateUserDetailsFirstName(t *testing.T) {
	db, mock, err := sqlmock.New()
	repo := NewUserDetailRepository(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	uuid := uuid.New().String()

	t.Run("Should update user detail first name", func(t *testing.T) {
		userDetail := UserDetail{Uuid: uuid, FirstName: "firstName", LastName: "lastName"}
		mock.ExpectExec(regexp.QuoteMeta("UPDATE UserDetails SET firstName=? WHERE uuid=?")).WithArgs(userDetail.FirstName, userDetail.Uuid).WillReturnResult(sqlmock.NewResult(0, 1))
		err = repo.UpdateFirstName(userDetail)

		if err != nil {
			t.Fatalf("User Details first name has not been updated, actual error: %v", err)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})

	t.Run("Should NOT update user detail first name when first name is not provided", func(t *testing.T) {
		userDetail := UserDetail{Uuid: uuid, LastName: "lastName"}
		err = repo.UpdateFirstName(userDetail)

		if err != err.(ErrMissingFirstName) {
			t.Fatalf("An error different than expected occured, actual error: %v", err)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})

	t.Run("Should NOT update user detail first name when uuid is not provided", func(t *testing.T) {
		userDetail := UserDetail{FirstName: "FirstName", LastName: "lastName"}
		err = repo.UpdateFirstName(userDetail)
		if err != err.(ErrMissingUuid) {
			t.Fatalf(fmt.Sprintf("An error different than expected occured, actual error: %v", err))
		}
	})
}

func TestUpdateUserDetailsLastName(t *testing.T) {
	db, mock, err := sqlmock.New()
	repo := NewUserDetailRepository(db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	uuid := uuid.New().String()

	t.Run("Should update user last name", func(t *testing.T) {
		userDetail := UserDetail{Uuid: uuid, FirstName: "firstName", LastName: "lastName"}
		mock.ExpectExec(regexp.QuoteMeta("UPDATE UserDetails SET lastName=? WHERE uuid=?")).WithArgs(userDetail.LastName, userDetail.Uuid).WillReturnResult(sqlmock.NewResult(0, 1))
		err = repo.UpdateLastName(userDetail)

		if err != nil {
			t.Fatalf("User Details last name has not been updated, actual error: %v", err)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})

	t.Run("Should NOT update user last name when last name is not provided", func(t *testing.T) {
		userDetail := UserDetail{Uuid: uuid, FirstName: "firstName"}
		err = repo.UpdateLastName(userDetail)

		if err != err.(ErrMissingLastName) {
			t.Fatalf("An error different than expected occured, actual error: %v", err)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})

	t.Run("Should NOT update user last name when uuid is not provided", func(t *testing.T) {
		userDetail := UserDetail{FirstName: "firstName", LastName: "lastName"}
		err = repo.UpdateLastName(userDetail)
		if err != err.(ErrMissingUuid) {
			t.Fatalf(fmt.Sprintf("An error different than expected occured, actual error: %v", err))
		}
	})
}
