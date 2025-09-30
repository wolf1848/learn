package validator

import "errors"

var ErrIsEmpty = errors.New("поле обязательно для заполнения")
var ErrIsMin = errors.New("поле должно содержать минимум %d символов")
var ErrIsMax = errors.New("поле не должно содержать больше %d символов")
var ErrIsEmail = errors.New("поле должно содержать валидный email адрес")
var ErrIsUnique = errors.New("поле должно быть уникальным")
var ErrInvalidValue = errors.New("не корректное значение")
