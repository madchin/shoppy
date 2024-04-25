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

func TestShouldGetUserDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	uuid := uuid.New().String()
	userDetail := &UserDetail{Uuid: uuid, FirstName: "firstName", LastName: "lastName"}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM UserDetails WHERE uuid=?")).WithArgs(uuid).WillReturnRows(sqlmock.NewRowsWithColumnDefinition(userDetailColumns...).AddRow(userDetail.Uuid, userDetail.FirstName, userDetail.LastName))
	selectedUserDetails, err := GetUserDetails(db, uuid)
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
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	uuid := uuid.New().String()
	userDetail := &UserDetail{Uuid: uuid, FirstName: "firstName", LastName: "lastName"}

	t.Run("Should create user details", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO UserDetails (uuid, firstName, lastName) VALUES (?, ?, ?)")).WithArgs(userDetail.Uuid, userDetail.FirstName, userDetail.LastName).WillReturnResult(sqlmock.NewResult(0, 1))
		err = userDetail.Create(db)

		if err != nil {
			t.Fatalf("User Details record has not been created, actual error: %v", err)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})

	t.Run("Should NOT create user details when uuid is not provided", func(t *testing.T) {
		userDetail.Uuid = ""
		err = userDetail.Create(db)

		if err != err.(*ErrMissingUuid) {
			t.Fatalf("An error different than expected occured, actual error: %v", err)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})

}

func TestUpdateUserDetailsFirstName(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	uuid := uuid.New().String()
	userDetail := &UserDetail{Uuid: uuid, FirstName: "firstName", LastName: "lastName"}

	t.Run("Should update user detail first name", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("UPDATE UserDetails SET firstName=? WHERE uuid=?")).WithArgs(userDetail.FirstName, userDetail.Uuid).WillReturnResult(sqlmock.NewResult(0, 1))
		err = userDetail.UpdateFirstName(db)

		if err != nil {
			t.Fatalf("User Details first name has not been updated, actual error: %v", err)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})

	t.Run("Should NOT update user detail first name when first name is not provided", func(t *testing.T) {
		userDetail.FirstName = ""
		err = userDetail.UpdateFirstName(db)

		if err != err.(*ErrMissingFirstName) {
			t.Fatalf("An error different than expected occured, actual error: %v", err)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})
}

func TestUpdateUserDetailsLastName(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	uuid := uuid.New().String()
	userDetail := &UserDetail{Uuid: uuid, FirstName: "firstName", LastName: "lastName"}

	t.Run("Should update user last name", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("UPDATE UserDetails SET lastName=? WHERE uuid=?")).WithArgs(userDetail.LastName, userDetail.Uuid).WillReturnResult(sqlmock.NewResult(0, 1))
		err = userDetail.UpdateLastName(db)

		if err != nil {
			t.Fatalf("User Details last name has not been updated, actual error: %v", err)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})

	t.Run("Should NOT update user last name when last name is not provided", func(t *testing.T) {
		userDetail.LastName = ""
		err = userDetail.UpdateLastName(db)

		if err != err.(*ErrMissingLastName) {
			t.Fatalf("An error different than expected occured, actual error: %v", err)
		}

		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})
}
