package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func GitPass(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "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",
	})
}

func Test1(c *gin.Context) {
	fmt.Println(c.ClientIP())
	fmt.Println(c.FullPath())

	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "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",
	})
}
