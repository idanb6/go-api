package main

import (
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
