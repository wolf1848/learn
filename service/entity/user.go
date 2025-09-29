package entity

import "errors"

var ErrNoRows = errors.New("нет данных соответствующих запросу")
var ErrAuthorized = errors.New("некорректные учетные данные")
var ErrInvalidToken = errors.New("токен не действителен")
var ErrService = errors.New("сервис временно не доступен, попробуйте позже")

type ErrValidRegister struct {
	Field string
	Err   error
}

func (e *ErrValidRegister) Error() string {
	return e.Err.Error()
}

func NewErrValidRegister(field string, err error) error {
	return &ErrValidRegister{
		Field: field,
		Err:   err,
	}
}

type RegisterInput struct {
	Name  string
	Email string
	Pwd   string
}

type RegisterOutput struct {
	ID    int
	Name  string
	Email string
}

type AuthorizeInput struct {
	Email string
	Pwd   string
}

type AuthorizeOutput struct {
	ID           int
	Name         string
	Email        string
	Token        string
	RefreshToken string
}
