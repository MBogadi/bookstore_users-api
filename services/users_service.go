package services

import (
	"../domain/users"
	"../utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil
}
