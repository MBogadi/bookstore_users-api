package app

import (
	gin "github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {

	mapUrls()
	router.Run(":8080")
}
