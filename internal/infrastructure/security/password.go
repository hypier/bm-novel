package security

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Hash 生成密码
func Hash(password string) ([]byte, error) {
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return fromPassword, nil
}

// VerifyPassword 验证密码
func VerifyPassword(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
