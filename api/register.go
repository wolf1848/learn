package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolf1848/taxiportal/api/dto"
	"github.com/wolf1848/taxiportal/service/entity"
)

type RegisterService interface {
	Register(input *entity.RegisterInput) (*entity.RegisterOutput, error)
}

func register(service RegisterService) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req *dto.RegisterRequest

		if err := c.Bind(&req); err != nil {
			return invalidJsonResponse(c)
		}

		var validateError = make(map[string][]string)

		if err, ok := notEmptyAndMinMaxValidate(req.Name, 3, 50); !ok {
			validateError["name"] = err
		}

		if err, ok := emailValidate(req.Email); !ok {
			validateError["email"] = err
		}

		if err, ok := notEmptyAndMinMaxValidate(req.Pwd, 8, 24); !ok {
			validateError["pwd"] = err
		}

		if len(validateError) == 0 {

			output, err := service.Register(&entity.RegisterInput{
				Name:  req.Name,
				Email: req.Email,
				Pwd:   req.Pwd,
			})

			if err == nil {
				response := &dto.RegisterResponse{
					ID:    output.ID,
					Name:  output.Name,
					Email: output.Email,
				}

				return c.JSON(http.StatusOK, response)
			}

			var errValid *entity.ErrValidRegister
			if errors.As(err, &errValid) {
				validateError[errValid.Field] = append(validateError[errValid.Field], errValid.Err.Error())
			}

			if errors.Is(err, entity.ErrService) {
				internalServerResponse(c)
			}

		}

		return invalidDataResponse(c, validateError)
	}
}
