package data

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

var UserColumns = []*sqlmock.Column{
	sqlmock.NewColumn("uuid").OfType("varchar(36)", uuid.NewString()).Nullable(false),
	sqlmock.NewColumn("name").OfType("varchar(255)", "randomName"),
	sqlmock.NewColumn("email").OfType("varchar(255)", "email@email.com").Nullable(false),
}
