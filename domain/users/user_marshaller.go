package users

import "encoding/json"

type PublicUser struct {
	Id          int64
	DateCreated string `json:"date_created"`
	Status      string
}

type PrivateUser struct {
	Id          int64
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string
	Status      string
	DateCreated string `json:"date_created"`
}

func (users Users) Marshall(public bool) []interface{} {
	result := make([]interface{}, len(users))
	for idx, user := range users {
		result[idx] = user.Marshall(public)
	}
	return result
}

func (user *User) Marshall(public bool) interface{} {
	if public {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}

	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)

	return privateUser
}
