package data

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

var columns = []*sqlmock.Column{
	sqlmock.NewColumn("uuid").OfType("varchar(36)", uuid.NewString()).Nullable(false),
	sqlmock.NewColumn("name").OfType("varchar(255)", "randomName"),
	sqlmock.NewColumn("email").OfType("varchar(255)", "email@email.com").Nullable(false),
}

func TestShouldGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	uuid := uuid.New().String()
	user := &User{Uuid: uuid, Name: "randomName", Email: "email@email.com"}

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM Users WHERE uuid=?")).WithArgs(uuid).WillReturnRows(sqlmock.NewRowsWithColumnDefinition(columns...).AddRow(user.Uuid, user.Name, user.Email))
	selectedUser, err := GetUser(db, uuid)
	if err != nil {
		t.Fatalf(fmt.Sprintf("%v", err))
	}

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error: %s", err)
	}
	if selectedUser.Name != user.Name {
		t.Fatalf("selected user name is not equal user inserted in db, expected: %s, actual: %s: ", user.Name, selectedUser.Name)
	}
	if selectedUser.Uuid != user.Uuid {
		t.Fatalf("selected user uuid is not equal user inserted in db, expected: %s, actual: %s: ", user.Uuid, selectedUser.Uuid)
	}
	if selectedUser.Email != user.Email {
		t.Fatalf("selected user email is not equal user inserted in db, expected: %s, actual: %s: ", user.Email, selectedUser.Email)
	}
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	uuid := uuid.New().String()
	user := &User{Uuid: uuid, Name: "randomName", Email: "email@email.com"}

	t.Run("should create user", func(t *testing.T) {
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO Users (uuid, name, email) VALUES (?, ?, ?)")).WithArgs(user.Uuid, user.Name, user.Email).WillReturnResult(sqlmock.NewResult(0, 1))
		err = user.Create(db)
		if err != nil {
			t.Fatalf(fmt.Sprintf("%v", err))
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("unmet expectation error: %s", err)
		}
	})

	t.Run("should NOT create when user email is empty", func(t *testing.T) {
		user.Email = ""
		err = user.Create(db)
		if err != ValidationError.EmptyEmail {
			t.Fatalf(fmt.Sprintf("%v", err))
		}
	})

}

func TestShouldUpdateName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	uuid := uuid.New().String()
	user := &User{Uuid: uuid, Name: "randomName", Email: "email@email.com"}

	sqlmock.NewRowsWithColumnDefinition(columns...).AddRow(user.Uuid, user.Name, user.Email)
	updateName := "updatedName"
	user.Name = updateName
	mock.ExpectExec(regexp.QuoteMeta("UPDATE Users SET name=? WHERE uuid=?")).WithArgs(user.Name, user.Uuid).WillReturnResult(sqlmock.NewResult(0, 1))

	err = user.UpdateName(db)
	if err != nil {
		t.Fatalf(fmt.Sprintf("%v", err))
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error: %s", err)
	}
	if user.Name != updateName {
		t.Fatalf("user name has not been updated, actual: %s, expected: %s", user.Name, updateName)
	}
}

func TestUpdateEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	uuid := uuid.New().String()
	user := &User{Uuid: uuid, Name: "randomName", Email: "email@email.com"}
	mock.ExpectExec(regexp.QuoteMeta("UPDATE Users SET email=? WHERE uuid=?")).WithArgs(user.Email, user.Uuid).WillReturnResult(sqlmock.NewResult(0, 1))
	err = user.UpdateEmail(db)
	if err != nil {
		t.Fatalf(fmt.Sprintf("%v", err))
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error: %s", err)
	}

}
