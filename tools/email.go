package tools

import "github.com/badoux/checkmail"

//EmailValidate Validate
func EmailValidate(email string) error {
	return checkmail.ValidateFormat(email)
}
