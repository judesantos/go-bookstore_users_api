package users

import (
	"github.com/judesantos/go-bookstore_users_api/utils/errors"
	"strings"
)

type User struct {
	Id          int64
	FirstName   string
	LastName    string
	Email       string
	DateCreated string
}

func (user *User) Validate() *errors.RestError {

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.InvalidParameterError("Invalid email address")
	}
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	if user.FirstName == "" {
		return errors.InvalidParameterError("Invalid Firstname")
	}
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	if user.LastName == "" {
		return errors.InvalidParameterError("Invalid Lastname")
	}

	return nil
}
