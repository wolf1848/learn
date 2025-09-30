package rules

import (
	"fmt"
	"github.com/wolf1848/taxiportal/validator"
)

const (
	nameMinLen = 3
	nameMaxLen = 50
)

func Name(value string) validator.ErrorList {
	errs := make(validator.ErrorList, 0)

	if validator.IsEmpty(value) {
		errs.Add(validator.NewError(validator.FieldName, ErrIsEmpty.Error()))
	}

	if validator.IsMinLen(value, nameMinLen) {
		errs.Add(validator.NewError(validator.FieldName, fmt.Sprintf(ErrIsMin.Error(), nameMinLen)))
	}

	if validator.IsMaxLen(value, nameMaxLen) {
		errs.Add(validator.NewError(validator.FieldName, fmt.Sprintf(ErrIsMax.Error(), nameMaxLen)))
	}

	return errs
}
