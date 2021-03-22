package tools

import "golang.org/x/crypto/bcrypt"

const cost = 10

//PasswordHash Hash Password
func PasswordHash(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}

//PasswordValidate Validate Password
func PasswordValidate(password string, hashed string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) != nil {
		return false
	}
	return true
}
