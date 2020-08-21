package services

import (
	"github.com/judesantos/go-bookstore_users_api/domain/users"
	crypto_utils "github.com/judesantos/go-bookstore_users_api/utils/crypto"
	"github.com/judesantos/go-bookstore_utils/rest_errors"
)

var (
	UsersService IUserService = &usersService{}
)

type usersService struct{}

type IUserService interface {
	CreateUser(users.User) (*users.User, rest_errors.IRestError)
	UpdateUser(bool, users.User) (*users.User, rest_errors.IRestError)
	DeleteUser(int64) rest_errors.IRestError
	GetUser(int64) (*users.User, rest_errors.IRestError)
	SearchUser(string) (users.Users, rest_errors.IRestError)
	LoginUser(users.LoginRequest) (*users.User, rest_errors.IRestError)
}

func (s *usersService) CreateUser(
	user users.User,
) (*users.User, rest_errors.IRestError) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) UpdateUser(
	partial bool,
	user users.User,
) (*users.User, rest_errors.IRestError) {

	_user, err := UsersService.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if partial {
		if user.FirstName != "" {
			_user.FirstName = user.FirstName
		}
		if user.LastName != "" {
			_user.LastName = user.LastName
		}
		if user.Email != "" {
			_user.Email = user.Email
		}
		if user.Status != "" {
			_user.Status = user.Status
		}
		if user.Password != "" {
			_user.Password = crypto_utils.GetMd5(user.Password)
		}
	} else {

		if err := user.Validate(); err != nil {
			return nil, err
		}

		_user.FirstName = user.FirstName
		_user.LastName = user.LastName
		_user.Email = user.Email
	}

	if err := _user.Update(); err != nil {
		return nil, err
	}

	return _user, nil
}

func (s *usersService) DeleteUser(userId int64) rest_errors.IRestError {
	_user, err := UsersService.GetUser(userId)
	if err != nil {
		return err
	}

	if err := _user.Delete(); err != nil {
		return err
	}

	return nil
}

func (s *usersService) GetUser(
	userId int64,
) (*users.User, rest_errors.IRestError) {

	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *usersService) SearchUser(
	status string,
) (users.Users, rest_errors.IRestError) {

	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(
	req users.LoginRequest,
) (*users.User, rest_errors.IRestError) {

	user := &users.User{
		Email:    req.Email,
		Password: req.Password,
	}

	if err := user.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	return user, nil
}
