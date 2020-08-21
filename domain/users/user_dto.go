package users

// user.dto.go

import (
	"strings"

	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

type User struct {
	Id          int64
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string
	Status      string
	Password    string
	DateCreated string `json:"date_created"`
}

type LoginRequest struct {
	Email    string
	Password string
}

type Users []User

func (user *User) Validate() rest_errors.IRestError {

	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.InvalidParameterError("Invalid email address")
	}

	return nil
}
