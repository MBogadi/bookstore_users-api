package users

import (
	"../../utils/errors"
	"fmt"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Email = result.Email
	user.FirstName = result.Email
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestError {

	if userDB[user.Id] != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	userDB[user.Id] = user
	return nil
}
