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

func FindByStatus(status string) ([]users.User, *errors.RestError) {
	var user users.User
	foundUsers, err := user.FindByStatus(status)

	if err != nil {
		return nil, err
	}

	// Return rows
	return foundUsers, nil
}

func UpdateUser(user *users.User) (*users.User, *errors.RestError) {
	currentUser, getError := GetUser(user.Id)
	if getError != nil {
		return nil, getError
	}

	if user.FirstName != "" {
		currentUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		currentUser.LastName = user.LastName
	}
	if user.Email != "" {
		currentUser.Email = user.Email
	}

	err := currentUser.Update()
	if err != nil {
		return nil, err
	}

	// updated
	return currentUser, nil
}

func DeleteUser(user *users.User) *errors.RestError {
	err := user.Delete()
	if err != nil {
		return err
	}

	// Deleted
	return nil
}
