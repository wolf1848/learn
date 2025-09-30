package rules

import (
	"github.com/wolf1848/taxiportal/validator"
)

func Email(value string) validator.ErrorList {
	errs := make(validator.ErrorList, 0)

	if validator.IsEmpty(value) {
		errs.Add(validator.NewError(validator.FieldEmail, ErrIsEmpty.Error()))
	}

	if !validator.IsEmail(value) {
		errs.Add(validator.NewError(validator.FieldEmail, ErrIsEmail.Error()))
	}

	return errs
}
