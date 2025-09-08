package service

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p []byte) (string, error) {
	cost := 14
	hash, err := bcrypt.GenerateFromPassword(p, cost)

	if err != nil {
		Logger().Error("Error on generate hash password")
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hashed []byte, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashed, password)

	if err != nil {
		Logger().Error("Error on compare hash password")
		return err
	}

	return nil
}
