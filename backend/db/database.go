package data

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	dbDriver "github.com/go-sql-driver/mysql"
)

func InitStore() (*sql.DB, error) {
	config := dbDriver.NewConfig()
	config.User = os.Getenv("DB_USER")
	config.Passwd = os.Getenv("DB_PASS")
	config.DBName = os.Getenv("DB_NAME")
	config.Addr = fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	var (
		db  *sql.DB
		err error
	)
	connector, err := dbDriver.NewConnector(config)
	db = sql.OpenDB(connector)

	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	seed(db)

	return db, nil
}

func seed(db *sql.DB) error {
	//FIXME workaround for initializing db
	time.Sleep(10 * time.Second)

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS Users (uuid varchar(36) NOT NULL, name varchar(255), email varchar(255) NOT NULL UNIQUE, PRIMARY KEY (uuid));")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS UserDetails (uuid varchar(36), firstName varchar(255) NOT NULL, lastName varchar(255) NOT NULL, PRIMARY KEY (uuid), FOREIGN KEY (uuid) REFERENCES Users(uuid) ON DELETE CASCADE);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Addresses (userId varchar(36), id int, postalCode varchar(64), address varchar(255), country varchar(128), city varchar(128), PRIMARY KEY (id), FOREIGN KEY (userId) REFERENCES Users(uuid) ON DELETE CASCADE);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Phones (userId varchar(36), id int, number varchar(255), PRIMARY KEY (id), FOREIGN KEY (userId) REFERENCES Users(uuid) ON DELETE CASCADE);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Products (id int, name varchar(255) UNIQUE, count int, price DECIMAL(5,2), PRIMARY KEY (id));")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Categories (id int, name varchar(255) UNIQUE, PRIMARY KEY (id));")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS ProductCategories (productId int, categoryId int, PRIMARY KEY (productId, categoryId), FOREIGN KEY (productId) REFERENCES Products(id) ON DELETE CASCADE, FOREIGN KEY (categoryId) REFERENCES Categories(id) ON DELETE CASCADE);")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Shippers (id int, name varchar(255) UNIQUE, PRIMARY KEY(id));")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Orders (id int, deliveryAddressId int, productId int, shipperId int, PRIMARY KEY(id), FOREIGN KEY(deliveryAddressId) REFERENCES Addresses(id), FOREIGN KEY(productId) REFERENCES Products(id), FOREIGN KEY(shipperId) REFERENCES Shippers(id));")
	if err != nil {
		return err
	}
	return nil
}
