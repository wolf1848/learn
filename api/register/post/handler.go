package post

import (
	"errors"
	"github.com/wolf1848/taxiportal/validator"
	"github.com/wolf1848/taxiportal/validator/rules"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolf1848/taxiportal/api/register/post/dto"
	"github.com/wolf1848/taxiportal/api/response"
	serviceErrors "github.com/wolf1848/taxiportal/service/errors"
	"github.com/wolf1848/taxiportal/service/register/entity"
)

type Service interface {
	Register(*entity.Input) (*entity.Output, error)
}

func Handler(service Service) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req *dto.Request

		if err := c.Bind(&req); err != nil {
			return response.InvalidJson(c)
		}

		validateErrs := validate(req)

		if len(validateErrs) > 0 {
			return response.InvalidDataRequest(c, validateErrs)
		}

		output, err := service.Register(&entity.Input{
			Name:  req.Name,
			Email: req.Email,
			Pwd:   req.Pwd,
		})

		if err != nil {
			if errors.Is(err, serviceErrors.ErrService) {
				return response.InternalServerError(c)
			}

			if errors.Is(err, entity.ErrUniqueEmail) {
				validateErrs.Add(validator.NewError(validator.FieldEmail, rules.ErrIsUnique.Error()))
			}

			if errors.Is(err, entity.ErrHashPwd) {
				validateErrs.Add(validator.NewError(validator.FieldPassword, rules.ErrInvalidValue.Error()))
			}

			return response.InvalidDataRequest(c, validateErrs)
		}

		responseData := &dto.Response{
			ID:    output.ID,
			Name:  output.Name,
			Email: output.Email,
		}

		return c.JSON(http.StatusOK, responseData)

	}
}
