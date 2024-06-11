package adapters

import (
	custom_error "backend/internal/common/errors"
	"backend/internal/users/domain/user"
	"database/sql"
)

type addressRepository struct {
	db *sql.DB
}

type addressDto struct {
	uuid       string
	street     string
	postalCode string
	country    string
	city       string
}

type addressesDto []addressDto

func NewAddressRepository(db *sql.DB) user.AddressRepository {
	return addressRepository{db}
}

func (a addressRepository) Create(userUuid string, address user.Address, validateFn func(user.Address) []error) custom_error.ContextError {
	if errs := validateFn(address); len(errs) > 0 {
		return custom_error.NewValidationErrors("address creation", errs)
	}
	_, err := a.db.Exec("INSERT INTO Addresses (uuid, postalCode, street, country, city) VALUES (?, ?, ?, ?, ?)", userUuid, address.PostalCode(), address.Street(), address.Country(), address.City())
	if err != nil {
		return custom_error.UnknownPersistenceError("address creation")
	}
	return custom_error.ContextError{}
}

// DeleteAddress implements user.AddressRepository.
func (a addressRepository) DeleteAddress(userUuid string, street string) custom_error.ContextError {
	addresses, err := a.Get(userUuid)
	if err.Error() != "" {
		return err
	}
	if !addresses.AddressExists(street) {
		return custom_error.NewPersistenceError("address delete", "address with provided street do not exists")
	}
	if _, err := a.db.Exec("DELETE * FROM Addresses WHERE uuid=? AND street=?", userUuid, street); err != nil {
		return custom_error.UnknownPersistenceError("address delete")
	}
	return custom_error.ContextError{}
}

// DeleteAll implements user.AddressRepository.
func (a addressRepository) DeleteAll(userUuid string) custom_error.ContextError {
	_, err := a.Get(userUuid)
	if err.Error() != "" {
		return err
	}
	if _, err := a.db.Exec("DELETE * FROM Addresses WHERE uuid=?", userUuid); err != nil {
		return custom_error.UnknownPersistenceError("address delete")
	}
	return custom_error.ContextError{}
}

func (a addressRepository) Get(userUuid string) (user.Addresses, custom_error.ContextError) {
	rows, err := a.db.Query("SELECT (uuid, postalCode, street, country, city) FROM Addresses WHERE uuid=?", userUuid)
	if err == sql.ErrNoRows {
		return user.Addresses{}, custom_error.NewPersistenceError("address retrieve", "user do not have any addresses")
	}
	if err != nil {
		return user.Addresses{}, custom_error.UnknownPersistenceError("address retrieve")
	}
	defer rows.Close()
	var addressesDto addressesDto
	for rows.Next() {
		addressDto := addressDto{}
		err = rows.Scan(addressDto.uuid, addressDto.postalCode, addressDto.street, addressDto.country, addressDto.city)
		if err != nil {
			addresses := addressesDto.mapToDomainAddresses()
			return addresses, custom_error.NewPersistenceError("address retrieve", "error during iterating address rows")
		}
		addressesDto = append(addressesDto, addressDto)
	}
	addresses := addressesDto.mapToDomainAddresses()
	return addresses, custom_error.ContextError{}
}

func (a addressRepository) Update(userUuid string, addressStreet string, address user.Address, validateFn func(user.Address) []error) custom_error.ContextError {
	if errs := validateFn(address); len(errs) > 0 {
		return custom_error.NewValidationErrors("address update", errs)
	}
	addresses, cerr := a.Get(userUuid)
	if cerr.Error() != "" {
		return cerr
	}
	if !addresses.AddressExists(addressStreet) {
		return custom_error.NewPersistenceError("address update", "address to update do not exists")
	}
	_, err := a.db.Exec("UPDATE Addresses SET city = ?, country = ?, postal_code = ?, street = ? WHERE uuid = ?", address.City(), address.Country(), address.PostalCode(), address.Street(), userUuid)
	if err != nil {
		return custom_error.UnknownPersistenceError("user address update")
	}
	return custom_error.ContextError{}
}

func (ad addressesDto) mapToDomainAddresses() user.Addresses {
	var addresses user.Addresses
	for _, addressDto := range ad {
		address := user.NewAddress(addressDto.postalCode, addressDto.street, addressDto.country, addressDto.city)
		addresses = append(addresses, address)
	}
	return addresses
}
