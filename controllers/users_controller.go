package controllers

import (
	"../domain/users"
	"../services"
	"../utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		//error handler
		restError := errors.NewBadRequestError("Invalid input JSON")
		c.JSON(restError.Status, restError)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//Error handler
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userid, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restError := errors.NewBadRequestError("Missing user_id in request")
		c.JSON(restError.Status, restError)
	}
	user, getError := services.GetUser(userid)
	if getError != nil {
		c.JSON(getError.Status, getError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userid, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restError := errors.NewBadRequestError("Missing user_id in request")
		c.JSON(restError.Status, restError)
	}

	var updateUser users.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		//error handler
		restError := errors.NewBadRequestError("Invalid input JSON")
		c.JSON(restError.Status, restError)
		return
	}

	updateUser.Id = userid
	updatedUser, updateError := services.UpdateUser(&updateUser)
	if updateError != nil {
		c.JSON(updateError.Status, updateError)
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c *gin.Context) {
	userid, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restError := errors.NewBadRequestError("Missing user_id in request")
		c.JSON(restError.Status, restError)
	}

	var deleteUser users.User
	deleteUser.Id = userid

	deleteError := services.DeleteUser(&deleteUser)
	if deleteError != nil {
		c.JSON(deleteError.Status, deleteError)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func FindByStatus(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		restError := errors.NewBadRequestError("Missing status in request")
		c.JSON(restError.Status, restError)
	}

	users, searchError := services.FindByStatus(status)
	if searchError != nil {
		c.JSON(searchError.Status, searchError)
		return
	}
	c.JSON(http.StatusOK, users)
}
