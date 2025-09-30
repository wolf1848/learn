package errors

import "errors"

var ErrService = errors.New("сервис временно не доступен, попробуйте позже")
var ErrRepositoryNoRows = errors.New("пустой ответ репозитория")
var ErrAuthorized = errors.New("некорректные учетные данные")
