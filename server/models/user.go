package models

import (
	"movieapp/errors"
	"regexp"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (user *User) ValidationUser() (validError []errors.ValidationError) {
	if user.FirstName == "" {
		validError = append(validError, errors.ValidationError{
			Fields:       "FirstName",
			MessageError: "FirstName is empty",
		})
	}
	if user.LastName == "" {
		validError = append(validError, errors.ValidationError{
			Fields:       "LastName",
			MessageError: "LastName is empty",
		})
	}
	if matched, err := regexp.Match(`^\w+@\w+\.\w+$`, []byte(user.Email)); err != nil || !matched {
		validError = append(validError, errors.ValidationError{
			Fields:       "Email",
			MessageError: "Email incorrect",
		})
	}
	if !checkUserPasswordIsStrongEnough(user) {
		validError = append(validError, errors.ValidationError{
			Fields:       "Password",
			MessageError: "Password incorrect",
		})
	}

	return
}

func checkUserPasswordIsStrongEnough(user *User) bool {
	if user.Password == "" {
		return false
	}
	return true
}
