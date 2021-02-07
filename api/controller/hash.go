package controller

import "golang.org/x/crypto/bcrypt"

//Hash hashes the password
func Hash(password string) (string, error) {
	hpassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hpassword), nil
}

//CompareHash compares two password returns the bool
func CompareHash(hpassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hpassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
