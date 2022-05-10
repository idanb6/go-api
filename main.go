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
	router.GET("/test1", Test1)
	port := os.Getenv("PORT")
	dominos := "dominos"
	fmt.Println(dominos)
	if port == "" {
		port = ":9999"
	}
	router.Run(port)
}
