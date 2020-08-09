package users

import (
	"fmt"
	"github.com/judesantos/go-bookstore_users_api/utils/errors"
)

var (
	usersDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := usersDb[user.Id]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user %d nof found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestError {
	exists := usersDb[user.Id]
	if exists == nil {
		return errors.BadRequestError(fmt.Sprintf("user %d exists", user.Id))
	}

	usersDb[user.Id] = user

	return nil
}