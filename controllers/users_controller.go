package controllers

import (
	"../domain/users"
	"../services"
	"../utils/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetUser(c *gin.Context) {
	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(err)
	fmt.Println(string(bytes))
	c.String(http.StatusNotImplemented, "Please Implement me!")
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
	fmt.Println(user)

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//Error handler
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Please Implement me!")
}
