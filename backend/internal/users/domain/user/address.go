package user

import (
	custom_error "backend/internal/common/errors"
	"errors"
)

type Address struct {
	postalCode string
	street     string
	country    string
	city       string
}

type Addresses []Address

func NewAddress(postalCode string, street string, country string, city string) Address {
	return Address{postalCode, street, country, city}
}

func (a Address) PostalCode() string {
	return a.postalCode
}
func (a Address) Street() string {
	return a.street
}
func (a Address) Country() string {
	return a.country
}
func (a Address) City() string {
	return a.city
}

func (adresses Addresses) AddressExists(street string) bool {
	for _, address := range adresses {
		if address.street == street {
			return true
		}
	}
	return false
}

func (a Address) validatePostalCode() (err error) {
	if a.postalCode == "" {
		err = errors.New("postal code is empty")
	}
	return
}

func (a Address) validateStreet() (err error) {
	if a.street == "" {
		err = errors.New("street is empty")
	}
	return
}

func (a Address) validateCountry() (err error) {
	if a.country == "" {
		err = errors.New("country is empty")
	}
	return
}

func (a Address) validateCity() (err error) {
	if a.city == "" {
		err = errors.New("city is empty")
	}
	return
}

func (a Address) Validate() (errs []error) {
	errs = custom_error.AppendError(errs, a.validateCity())
	errs = custom_error.AppendError(errs, a.validateCountry())
	errs = custom_error.AppendError(errs, a.validatePostalCode())
	errs = custom_error.AppendError(errs, a.validateStreet())

	return
}
