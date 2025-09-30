package post

import (
	"github.com/wolf1848/taxiportal/api/authorize/post/dto"
	"github.com/wolf1848/taxiportal/validator"
	"github.com/wolf1848/taxiportal/validator/rules"
)

func validate(req *dto.Request) validator.ErrorList {
	valid := validator.NewValidator().
		Add(rules.Email(req.Email)...).
		Add(rules.Password(req.Pwd)...)

	return valid.Validate()
}
