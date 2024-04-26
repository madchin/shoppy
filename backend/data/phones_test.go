package data

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

var phonesColumns = []*sqlmock.Column{
	sqlmock.NewColumn("uuid").OfType("varchar(36)", uuid.NewString()).Nullable(false),
	sqlmock.NewColumn("id").OfType("int", 1).Nullable(false),
	sqlmock.NewColumn("number").OfType("varchar(255)", "123123123123").Nullable(false),
}

func TestGetPhones(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	t.Run("Should get 2 phones", func(t *testing.T) {
		uuid := uuid.NewString()
		phone1 := &Phone{Uuid: uuid, Number: "123123123", Id: 1}
		phone2 := &Phone{Uuid: uuid, Number: "234234234", Id: 2}
		mock.ExpectQuery(regexp.QuoteMeta("SELECT id, number FROM Phones WHERE uuid=?")).WithArgs(uuid).WillReturnRows(sqlmock.NewRowsWithColumnDefinition(phonesColumns...).AddRow(phone1.Uuid, phone1.Id, phone1.Number).AddRow(phone2.Uuid, phone2.Id, phone2.Number))
		phones, err := GetPhones(db, uuid)
		if err != nil {
			t.Fatalf("An unexpected error occured, actual: %v", err)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("an unexpected expectations occured, actual: %v", err)
		}
		if len := len(phones); len < 2 {
			t.Fatalf("actual count of phones retrieved is %d, should be 2", len)
		}
		if phones[0].Uuid != phone1.Uuid {
			t.Fatalf("Retrieved first phone uuid is incorrect, retrieved uuid: %s, actual uuid: %s", phones[0].Uuid, phone1.Uuid)
		}
		if phones[1].Uuid != phone2.Uuid {
			t.Fatalf("Retrieved second phone uuid is incorrect, retrieved uuid: %s, actual uuid: %s", phones[1].Uuid, phone2.Uuid)
		}
		if phones[1].Number != phone2.Number {
			t.Fatalf("Retrieved second phone number is incorrect, retrieved number: %s, actual number: %s", phones[1].Number, phone2.Number)
		}
		if phones[0].Number != phone1.Number {
			t.Fatalf("Retrieved first phone number is incorrect, retrieved number: %s, actual number: %s", phones[0].Number, phone1.Number)
		}
		if phones[0].Id != phone1.Id {
			t.Fatalf("Retrieved first phone id is incorrect, retrieved id: %d, actual id: %d", phones[0].Id, phone1.Id)
		}
		if phones[1].Id != phone2.Id {
			t.Fatalf("Retrieved second phone id is incorrect, retrieved id: %d, actual id: %d", phones[1].Id, phone2.Id)
		}
	})
	t.Run("Should not get phone when uuid is empty", func(t *testing.T) {
		_, err := GetPhones(db, "")
		if err != err.(*ErrMissingUuid) {
			t.Fatalf("An unexpected error occured, actual error: %v", err)
		}
	})
}

func TestCreatePhone(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("Should NOT create phone when uuid is empty", func(t *testing.T) {
		p := &Phone{Number: "123123123"}
		err := p.Create(db)
		if err != err.(*ErrMissingUuid) {
			t.Fatalf("An unexpected error occured, actual: %v", err)
		}
	})

	t.Run("Should create phone", func(t *testing.T) {
		p := &Phone{Uuid: uuid.NewString(), Number: "12313223"}
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO Phones (uuid,number) VALUES (?,?)")).WithArgs(p.Uuid, p.Number).WillReturnResult(sqlmock.NewResult(1, 1))
		err := p.Create(db)
		if err != nil {
			t.Fatalf("An unexpected error occured when creating phone number, actual err: %v", err)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("An unexpected error occured, actual %v", err)
		}
	})
}

func TestUpdateNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("Should update number", func(t *testing.T) {
		p := &Phone{Uuid: uuid.NewString(), Number: "12313223"}
		mock.ExpectExec(regexp.QuoteMeta("UPDATE Phones SET number=? WHERE uuid=?")).WithArgs(p.Number, p.Uuid).WillReturnResult(sqlmock.NewResult(0, 1))
		err := p.Update(db)

		if err != nil {
			t.Fatalf("An unexpected error occured when updating phone number, actual err: %v", err)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("An unexpected error occured, actual %v", err)
		}
	})

	t.Run("Should NOT update number when number is empty", func(t *testing.T) {
		p := &Phone{Uuid: uuid.NewString()}
		err := p.Update(db)
		if err != err.(*ErrMissingPhoneNumber) {
			t.Fatalf("An unexpected error occured, actual: %v", err)
		}
	})
	t.Run("Should NOT update number when uuid is empty", func(t *testing.T) {
		p := &Phone{Number: "123132123"}
		err := p.Update(db)
		if err != err.(*ErrMissingUuid) {
			t.Fatalf("An unexpected error occured, actual: %v", err)
		}
	})
}

func TestDeleteNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	t.Run("Should delete number", func(t *testing.T) {
		p := &Phone{Uuid: uuid.NewString(), Number: "12313223"}
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM Phones WHERE uuid=? AND number=?")).WithArgs(p.Uuid, p.Number).WillReturnResult(sqlmock.NewResult(0, 1))
		err := p.Delete(db)
		if err != nil {
			t.Fatalf("An unexpected error occured when creating phone number, actual err: %v", err)
		}
		if err = mock.ExpectationsWereMet(); err != nil {
			t.Fatalf("An unexpected error occured, actual %v", err)
		}
	})
	t.Run("Should NOT delete number when number is empty", func(t *testing.T) {
		p := &Phone{Uuid: uuid.NewString()}
		err := p.Delete(db)
		if err != err.(*ErrMissingPhoneNumber) {
			t.Fatalf("An unexpected error occured, actual: %v", err)
		}
	})
	t.Run("Should NOT delete number when uuid is empty", func(t *testing.T) {
		p := &Phone{Number: "123132123"}
		err := p.Delete(db)
		if err != err.(*ErrMissingUuid) {
			t.Fatalf("An unexpected error occured, actual: %v", err)
		}
	})
}
