package controllers

import (
	"../domain/users"
	"../services"
	"../utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

func CreateUser(c *gin.Context) {
	var user users.User
	/*
		bytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			//Error handler
			fmt.Println("IO ReadAll error: ", err)
			return
		}

		if err := json.Unmarshal(bytes, &user); err != nil {
			//Error handler
			fmt.Println("JSON Unmarshal error: ", err)
			return
		}
		fmt.Printf("Bytes: %v", string(bytes))
		fmt.Println("")
	*/

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
