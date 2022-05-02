package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", GetHello)
	router.Run("localhost:9999")
}
