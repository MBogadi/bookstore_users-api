package services

import (
	"../domain/users"
	"../utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	//validate
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userID int64) (*users.User, *errors.RestError) {
	var user users.User
	user.Id = userID
	err := user.Get()

	if err != nil {
		return nil, err
	}

	// Found
	return &user, nil
}
