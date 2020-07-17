package BPT

import "golang.org/x/crypto/bcrypt"

func GenerateFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	newPassword := string(hash)
	return newPassword, nil
}


func CompareHashAndPassword(password,realPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(realPassword), []byte(password))
	if err != nil {
		return false
	}

	return true
}