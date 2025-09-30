package post

import (
	"fmt"

	"github.com/wolf1848/taxiportal/api/register/post/dto"
	"github.com/wolf1848/taxiportal/validator"
)

const (
	NAME_MIN_LEN = 3
	NAME_MAX_LEN = 50
	PWD_MIN_LEN = 8
	PWD_MAX_LEN = 24
)
func validate(data *dto.Request) *validator.Validator {
	valid := validator.NewValidator()

	if validator.IsEmpty(data.Name) {
		valid.AddProblem("name", validator.ErrIsEmpty.Error())
	}

	if validator.IsMinLen(data.Name, NAME_MIN_LEN) {
		valid.AddProblem("name", fmt.Sprintf(validator.ErrIsMin.Error(), NAME_MIN_LEN))
	}

	if validator.IsMaxLen(data.Name, NAME_MAX_LEN) {
		valid.AddProblem("name", fmt.Sprintf(validator.ErrIsMax.Error(), NAME_MAX_LEN))
	}

	if validator.IsEmpty(data.Email) {
		valid.AddProblem("email", validator.ErrIsEmpty.Error())
	}

	if !validator.IsEmail(data.Email) {
		valid.AddProblem("email", validator.ErrIsEmail.Error())
	}

	if validator.IsEmpty(data.Pwd) {
		valid.AddProblem("pwd", validator.ErrIsEmpty.Error())
	}

	if validator.IsMinLen(data.Pwd, PWD_MIN_LEN) {
		valid.AddProblem("pwd", fmt.Sprintf(validator.ErrIsMin.Error(), PWD_MIN_LEN))
	}

	if validator.IsMaxLen(data.Pwd, PWD_MAX_LEN) {
		valid.AddProblem("pwd", fmt.Sprintf(validator.ErrIsMax.Error(), PWD_MAX_LEN))
	}

	return valid
}
