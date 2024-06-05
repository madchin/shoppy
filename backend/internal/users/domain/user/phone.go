package user

import (
	"errors"
	"regexp"
)

type Phone struct {
	number string
}

type Phones []Phone

const (
	maxNumberLength = 15
	numberRegex     = `^[+]{1}(?:[0-9\-\(\)\/\.]\s?){6, 15}[0-9]{1}$`
)

func NewPhone(number string) Phone {
	return Phone{number}
}

func (phone Phone) Number() string {
	return phone.number
}

func (phones Phones) NumberExist(number string) bool {
	deletePhoneExists := false
	for _, phone := range phones {
		if phone.Number() == number {
			deletePhoneExists = true
			break
		}
	}
	return deletePhoneExists
}

func (phones Phones) AllPhoneNumbers() []string {
	var phoneNumbers []string
	for _, phone := range phones {
		phoneNumbers = append(phoneNumbers, phone.number)
	}
	return phoneNumbers
}

func (phone Phone) Validate() (errs []error) {
	return phone.validateNumber()
}

func (p Phone) validateNumber() (errs []error) {
	if len(p.number) > maxNumberLength {
		errs = append(errs, errors.New("Phone number is too long"))
	}
	ok, _ := regexp.MatchString(numberRegex, p.number)
	if !ok {
		errs = append(errs, errors.New("Phone format is wrong"))
	}
	return
}
