package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

type Execise struct {
	BodyPart  string `json:"bodyPart"`
	Equipment string `json:"equipment"`
	GifUrl    string `json:"gifUrl"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Target    string `json:"target"`
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", GetHello)
	router.GET("/git", GitPass)
	router.GET("/test1", Test1)
	router.GET("/execises", GetExecise)
	router.GET("/execise/:id", GetByIdExecise)
	port := os.Getenv("PORT")
	dominos := "dominos"
	fmt.Println(dominos)
	if port == "" {
		port = "9999"
	}
	router.Run(":" + port)
}
