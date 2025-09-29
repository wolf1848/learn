package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wolf1848/taxiportal/validator"
)

type invalidAuthorize struct {
	Message string `json:"message"`
}

var invalidAuthorizeData = &invalidJson{
	Message: "Некорректные учетные данные",
}

func invalidAuthorizeResponse(c echo.Context) error {
	return c.JSON(http.StatusUnauthorized, invalidAuthorizeData)
}

type invalidJson struct {
	Message string `json:"message"`
}

var invalidJsonData = &invalidJson{
	Message: "Некорректный JSON",
}

func invalidJsonResponse(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, invalidJsonData)
}

type internalServer struct {
	Message string `json:"message"`
}

var internalServerData = &internalServer{
	Message: "Сервер временно недоступен, попробуйте позже",
}

func internalServerResponse(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, internalServerData)
}

func invalidDataResponse(c echo.Context, errors any) error {
	return c.JSON(http.StatusUnprocessableEntity, errors)
}

func notEmptyAndMinMaxValidate(value string, min, max int) ([]string, bool) {
	var errors []string
	if validator.IsEmpty(value) {
		errors = append(errors, ErrRequire.Error())
	}

	if validator.IsMinLen(value, min) {
		errors = append(errors, fmt.Sprintf(ErrMin.Error(), min))
	}

	if validator.IsMaxLen(value, max) {
		errors = append(errors, fmt.Sprintf(ErrMax.Error(), max))
	}

	return errors, len(errors) == 0
}

func emailValidate(value string) ([]string, bool) {
	var errors []string

	if validator.IsEmpty(value) {
		errors = append(errors, ErrRequire.Error())
	}

	if !validator.IsEmail(value) {
		errors = append(errors, ErrEmail.Error())
	}

	return errors, len(errors) == 0
}
