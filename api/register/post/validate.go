package post

import (
	"github.com/wolf1848/taxiportal/api/register/post/dto"
	"github.com/wolf1848/taxiportal/validator"
	"github.com/wolf1848/taxiportal/validator/rules"
)

func validate(req *dto.Request) validator.ErrorList {
	valid := validator.NewValidator().
		Add(rules.Name(req.Name)...).
		Add(rules.Email(req.Email)...).
		Add(rules.Password(req.Pwd)...)

	return valid.Validate()
}
