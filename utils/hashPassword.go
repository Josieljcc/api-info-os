package utils

import "golang.org/x/crypto/bcrypt"

const salt = "infoOskeysalt"

func HashPassword(password string) (string, error) {
	saltedPassword := []byte(password + salt)
	hashedPassword, err := bcrypt.GenerateFromPassword(saltedPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) bool {
	saltedPassword := []byte(password + salt)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), saltedPassword)
	return err == nil
}
