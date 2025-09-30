package post

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolf1848/taxiportal/api/register/post/dto"
	"github.com/wolf1848/taxiportal/api/response"
	serviceErrors "github.com/wolf1848/taxiportal/service/errors"
	"github.com/wolf1848/taxiportal/service/register/entity"
	"github.com/wolf1848/taxiportal/validator"
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

		validateResult := validate(req)

		if !validateResult.IsValid() {
			return response.InvalidDataRequest(c, validateResult.GetProblems())
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
				validateResult.AddProblem("email", validator.ErrIsUnique.Error())
			}
			

			if errors.Is(err, entity.ErrHashPwd) {
				validateResult.AddProblem("pwd", validator.ErrInvalidValue.Error())
			}

			return response.InvalidDataRequest(c, validateResult.GetProblems())
		}

		responseData := &dto.Response{
			ID:    output.ID,
			Name:  output.Name,
			Email: output.Email,
		}

		return c.JSON(http.StatusOK, responseData)

	}
}
