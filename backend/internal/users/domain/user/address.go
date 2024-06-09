package user

import "errors"

type Address struct {
	postalCode string
	street     string
	country    string
	city       string
}

const ()

func NewAddress(postalCode string, street string, country string, city string) Address {
	return Address{postalCode, street, country, city}
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
	errs = append(errs, a.validateCity())
	errs = append(errs, a.validateCountry())
	errs = append(errs, a.validatePostalCode())
	errs = append(errs, a.validateStreet())

	return
}
