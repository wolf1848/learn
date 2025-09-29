package api

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolf1848/taxiportal/api/dto"
	"github.com/wolf1848/taxiportal/service/entity"
)

type AuthorizeService interface {
	Authorize(input *entity.AuthorizeInput) (*entity.AuthorizeOutput, error)
}

func authorize(service AuthorizeService) echo.HandlerFunc {
	return func(c echo.Context) error {

		var req *dto.AuthorizeRequest

		if err := c.Bind(&req); err != nil {
			return invalidJsonResponse(c)
		}

		var validateError = make(map[string][]string)

		if err, ok := emailValidate(req.Email); !ok {
			validateError["email"] = err
		}

		if err, ok := notEmptyAndMinMaxValidate(req.Pwd, 8, 24); !ok {
			validateError["pwd"] = err
		}

		if len(validateError) == 0 {

			output, err := service.Authorize(&entity.AuthorizeInput{
				Email: req.Email,
				Pwd:   req.Pwd,
			})

			if err != nil {
				if errors.Is(err, entity.ErrAuthorized) {
					return invalidAuthorizeResponse(c)
				}
				return internalServerResponse(c)
			}

			response := &dto.AuthorizeResponse{
				ID:           output.ID,
				Name:         output.Name,
				Email:        output.Email,
				Token:        output.Token,
				RefreshToken: output.RefreshToken,
			}

			return c.JSON(http.StatusOK, response)

		}

		return invalidDataResponse(c, validateError)
	}
}
