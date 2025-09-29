package api

import "errors"

var ErrRequire = errors.New("поле обязательно для заполнения")
var ErrMin = errors.New("поле должно содержать минимум %d символов")
var ErrMax = errors.New("поле не должно содержать больше %d символов")
var ErrEmail = errors.New("поле должно содержать валидный email адрес")
var ErrEmailUnique = errors.New("email адрес уже используется")
