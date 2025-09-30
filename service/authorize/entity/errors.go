package entity

import "errors"

var ErrUniqueEmail = errors.New("email уже используется")
var ErrHashPwd = errors.New("попробуйте изменить пароль")
