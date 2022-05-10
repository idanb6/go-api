package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", GetHello)
	router.GET("/git", GitPass)
	port := os.Getenv("PORT")
	fmt.Println("dominos")
	if port == "" {
		port = ":9999"
	}
	router.Run(port)
}
