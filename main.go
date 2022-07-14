package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/git", GitPass)
	router.GET("/test1", Test1)
	router.GET("/execises", GetExecise)
	router.GET("/execise/:id", GetByIdExecise)

	router.GET("/bodyPart/:bodyPart", GetByBodyPartExecise)
	router.GET("/bodyPart", GetBodyPart)

	router.GET("/equipment/:equipment", GetByEquipmentExecise)
	router.GET("/equipment", GetEquipment)

	router.GET("/target/:target", GetByTargetExecise)
	router.GET("/target", GetTarget)

	router.GET("/serch", GetByNameExecise)

	port := os.Getenv("PORT")
	dominos := "dominos"
	fmt.Println(dominos)
	if port == "" {
		port = "9999"
	}
	router.Run(":" + port)
}
