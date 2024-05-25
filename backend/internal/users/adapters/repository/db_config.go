package adapters

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var REQUIRED_ENVS = []string{"DB_USER", "DB_PASS", "DB_HOST", "DB_PORT", "DB_NAME"}

func NewDatabase() (*sql.DB, error) {
	missingEnvs := make(map[string]string, 0)
	for _, key := range REQUIRED_ENVS {
		if value, ok := os.LookupEnv(key); !ok {
			missingEnvs[key] = value
			continue
		}
	}
	if len(missingEnvs) >= 1 {
		panic(fmt.Sprintf("DB CONFIG ERROR: Missing envs: %v", missingEnvs))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("DB CONFIG ERROR: db has not been opened: %s", err.Error()))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = createTables(db)
	if err != nil {
		panic(fmt.Sprintf("DB CONFIG ERROR: tables has not been created: %s", err.Error()))
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
