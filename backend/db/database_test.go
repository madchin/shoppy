package db

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateTables(t *testing.T) {
	db, mock, err := sqlmock.New()
	mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS Users (uuid varchar(36) NOT NULL, name varchar(255), email varchar(255) NOT NULL UNIQUE, PRIMARY KEY (uuid))")).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS UserDetails (uuid varchar(36), firstName varchar(255) NOT NULL, lastName varchar(255) NOT NULL, PRIMARY KEY (uuid), FOREIGN KEY (uuid) REFERENCES Users(uuid) ON DELETE CASCADE)")).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS Addresses (uuid varchar(36), id int AUTO_INCREMENT, postalCode varchar(64), address varchar(255), country varchar(128), city varchar(128), PRIMARY KEY (id), FOREIGN KEY (uuid) REFERENCES Users(uuid) ON DELETE CASCADE)")).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS Phones (uuid varchar(36), id int AUTO_INCREMENT, number varchar(255) NOT NULL, PRIMARY KEY (id), FOREIGN KEY (uuid) REFERENCES Users(uuid) ON DELETE CASCADE)")).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS Products (id int AUTO_INCREMENT, name varchar(255) UNIQUE, count int, price DECIMAL(5,2), PRIMARY KEY (id))")).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS Categories (id int AUTO_INCREMENT, name varchar(255) UNIQUE, PRIMARY KEY (id))")).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS ProductCategories (productId int, categoryId int, PRIMARY KEY (productId, categoryId), FOREIGN KEY (productId) REFERENCES Products(id) ON DELETE CASCADE, FOREIGN KEY (categoryId) REFERENCES Categories(id) ON DELETE CASCADE)")).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS Shippers (id int AUTO_INCREMENT, name varchar(255) UNIQUE, PRIMARY KEY(id))")).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS Countries (id int AUTO_INCREMENT, name varchar(255) UNIQUE, PRIMARY KEY(id))")).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS ShipperCountries (shipperId int, countryId int, PRIMARY KEY(shipperId, countryId), FOREIGN KEY (shipperId) REFERENCES Shippers(id), FOREIGN KEY(countryId) REFERENCES Countries(id))")).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(regexp.QuoteMeta("CREATE TABLE IF NOT EXISTS Orders (id int AUTO_INCREMENT, deliveryAddressId int, productId int, shipperId int, PRIMARY KEY(id), FOREIGN KEY(deliveryAddressId) REFERENCES Addresses(id), FOREIGN KEY(productId) REFERENCES Products(id), FOREIGN KEY(shipperId) REFERENCES Shippers(id))")).WillReturnResult(sqlmock.NewResult(0, 1))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	err = createTables(db)
	if err != nil {
		t.Fatalf("an error occured creating table %v", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectation error: %s", err)
	}
}
