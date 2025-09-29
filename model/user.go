package model

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrUniqueEmail = errors.New("email уже используется")
var ErrHashPwd = errors.New("попробуйте изменить пароль")

type User struct {
	ID      int
	Name    string
	Email   string
	HashPwd string
}

func (model *User) SetPwd(pwd string) error {
	hash, err := hashPwd(pwd)
	if err != nil {
		return err
	}
	model.HashPwd = string(hash)
	return nil
}

func hashPwd(pwd string) (string, error) {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPwd), nil
}

func (model *User) CheckPasswordHash(pwd string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(model.HashPwd), []byte(pwd))
    return err == nil
}
