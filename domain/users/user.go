package users

import (
	"../../utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	user.Password = strings.TrimSpace(user.Password)
	user.Status = strings.TrimSpace(strings.ToLower(user.Status))

	if user.Status == "" {
		return errors.NewBadRequestError("Invalid user Status")
	}
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email ID")
	}
	if user.Password == "" {
		return errors.NewBadRequestError("Password cannot be blank")
	}

	return nil
}
