package adapters

import (
	"backend/internal/common/repository"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var REQUIRED_ENVS = []string{"DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME"}

func NewDatabase() (*sql.DB, error) {
	errMissingEnvs := &repository.ErrMissingEnv{}
	envs := make(map[string]string, len(REQUIRED_ENVS))
	for _, key := range REQUIRED_ENVS {
		if value, ok := os.LookupEnv(key); ok {
			envs[key] = value
			continue
		}
		errMissingEnvs.Add(key)
	}
	if len(errMissingEnvs.Keys) >= 1 {
		return nil, errMissingEnvs
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", envs["DB_USER"], envs["DB_PASS"], envs["DB_HOST"], envs["DB_PORT"], envs["DB_NAME"])
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS Users (uuid varchar(36) NOT NULL, name varchar(255), email varchar(255) NOT NULL UNIQUE, PRIMARY KEY (uuid))")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS UserDetails (uuid varchar(36), firstName varchar(255) NOT NULL, lastName varchar(255) NOT NULL, PRIMARY KEY (uuid), FOREIGN KEY (uuid) REFERENCES Users(uuid) ON DELETE CASCADE)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Addresses (uuid varchar(36), id int AUTO_INCREMENT, postalCode varchar(64), address varchar(255), country varchar(128), city varchar(128), PRIMARY KEY (id), FOREIGN KEY (uuid) REFERENCES Users(uuid) ON DELETE CASCADE)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Phones (uuid varchar(36), id int AUTO_INCREMENT, number varchar(255) NOT NULL, PRIMARY KEY (id), FOREIGN KEY (uuid) REFERENCES Users(uuid) ON DELETE CASCADE)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Products (id int AUTO_INCREMENT, name varchar(255) UNIQUE, count int, price DECIMAL(5,2), PRIMARY KEY (id))")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Categories (id int AUTO_INCREMENT, name varchar(255) UNIQUE, PRIMARY KEY (id))")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS ProductCategories (productId int, categoryId int, PRIMARY KEY (productId, categoryId), FOREIGN KEY (productId) REFERENCES Products(id) ON DELETE CASCADE, FOREIGN KEY (categoryId) REFERENCES Categories(id) ON DELETE CASCADE)")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Shippers (id int AUTO_INCREMENT, name varchar(255) UNIQUE, PRIMARY KEY(id))")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Countries (id int AUTO_INCREMENT, name varchar(255) UNIQUE, PRIMARY KEY(id))")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS ShipperCountries (shipperId int, countryId int, PRIMARY KEY(shipperId, countryId), FOREIGN KEY (shipperId) REFERENCES Shippers(id), FOREIGN KEY(countryId) REFERENCES Countries(id))")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Orders (id int AUTO_INCREMENT, deliveryAddressId int, productId int, shipperId int, PRIMARY KEY(id), FOREIGN KEY(deliveryAddressId) REFERENCES Addresses(id), FOREIGN KEY(productId) REFERENCES Products(id), FOREIGN KEY(shipperId) REFERENCES Shippers(id))")
	if err != nil {
		return err
	}

	return nil
}
