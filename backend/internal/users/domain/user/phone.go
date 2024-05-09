package user

import (
	"errors"
	"fmt"
	"regexp"

	cerr "backend/internal/common/error"

	"github.com/hashicorp/go-multierror"
)

type Phone struct {
	number int
}

var errNumberMaxLength = errors.New("Phone number is too long")
var errNumberNotMatch = errors.New("Phone format is wrong")

const (
	maxNumberLength = 15
	numberRegex     = `^[+]{1}(?:[0-9\-\(\)\/\.]\s?){6, 15}[0-9]{1}$`
)

func (p Phone) Validate() error {
	err := p.validateNumber()
	return err.(*multierror.Error).ErrorOrNil()
}

func (p Phone) IsProvided() bool {
	return p.Validate() != nil
}

func (p Phone) validateNumber() (err error) {
	strNumber := fmt.Sprintf("%d", p.number)
	if len(strNumber) > maxNumberLength {
		err = errNumberMaxLength
	}
	ok, rerr := regexp.MatchString(numberRegex, strNumber)
	if rerr != nil {
		err = multierror.Append(err, cerr.ErrInternal)
	}
	if !ok {
		err = multierror.Append(err, errNumberNotMatch)
	}
	return err
}
