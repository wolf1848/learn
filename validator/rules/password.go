package rules

import (
	"fmt"
	"github.com/wolf1848/taxiportal/validator"
)

const (
	pwdMinLen = 8
	pwdMaxLen = 24
)

func Password(password string) validator.ErrorList {
	errs := make(validator.ErrorList, 0)

	if validator.IsEmpty(password) {
		errs.Add(validator.NewError(validator.FieldPassword, ErrIsEmpty.Error()))
	}

	if validator.IsMinLen(password, pwdMinLen) {
		errs.Add(validator.NewError(validator.FieldPassword, fmt.Sprintf(ErrIsMin.Error(), pwdMinLen)))
	}

	if validator.IsMaxLen(password, pwdMaxLen) {
		errs.Add(validator.NewError(validator.FieldPassword, fmt.Sprintf(ErrIsMax.Error(), pwdMaxLen)))
	}

	return errs
}
