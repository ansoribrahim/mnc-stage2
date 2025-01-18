package util

import "golang.org/x/crypto/bcrypt"

func HashPin(pin string) (string, error) {
	hashedPin, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPin), nil
}

func VerifyPin(hashedPin, plainPin string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPin), []byte(plainPin))
	return err == nil
}
