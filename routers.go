package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetExecise(c *gin.Context) {
	ex, err := getExecise()
	if err != nil {
		c.JSON(400, gin.H{
			"data":    "error",
			"message": err,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "ok",
			"data":    ex,
		})
	}

}
func GetByIdExecise(c *gin.Context) {
	exs, err := getExecise()
	if err != nil {
		c.JSON(400, gin.H{
			"data":    "error",
			"message": err,
		})
	}
	id := c.Param("id")
	ex := findById(id, exs)
	if ex.Id == "" {
		c.JSON(200, gin.H{
			"message": "ok",
			"data":    nil,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "ok",
			"data":    ex,
		})
	}

}

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
