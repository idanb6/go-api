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
			"total":   len(ex),
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

func GetByNameExecise(c *gin.Context) {
	exs, err := getExecise()
	if err != nil {
		c.JSON(400, gin.H{
			"data":    "error",
			"message": err,
		})
	}
	text, ok := c.GetQuery("q")
	if !ok {
		c.JSON(400, gin.H{})
	}
	ex := findByType("serch", text, exs)

	c.JSON(200, gin.H{
		"message": "ok",
		"total":   len(ex),
		"data":    ex,
	})

}

func GetByTargetExecise(c *gin.Context) {
	exs, err := getExecise()
	if err != nil {
		c.JSON(400, gin.H{
			"data":    "error",
			"message": err,
		})
	}
	target := c.Param("target")
	ex := findByType("target", target, exs)
	if len(ex) == 0 {
		c.JSON(200, gin.H{
			"message": "ok",
			"data":    nil,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "ok",
			"total":   len(ex),
			"data":    ex,
		})
	}
}

func GetByEquipmentExecise(c *gin.Context) {
	exs, err := getExecise()
	if err != nil {
		c.JSON(400, gin.H{
			"data":    "error",
			"message": err,
		})
	}
	equipment := c.Param("equipment")
	ex := findByType("equipment", equipment, exs)
	if len(ex) == 0 {
		c.JSON(200, gin.H{
			"message": "ok",
			"data":    nil,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "ok",
			"total":   len(ex),
			"data":    ex,
		})
	}

}

func GetByBodyPartExecise(c *gin.Context) {
	exs, err := getExecise()
	if err != nil {
		c.JSON(400, gin.H{
			"data":    "error",
			"message": err,
		})
	}
	bodyPart := c.Param("bodyPart")
	ex := findByType("bodyPart", bodyPart, exs)
	if len(ex) == 0 {
		c.JSON(200, gin.H{
			"message": "ok",
			"data":    nil,
		})
	} else {
		c.JSON(200, gin.H{
			"message": "ok",
			"total":   len(ex),
			"data":    ex,
		})
	}

}

func GetBodyPart(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
		"data":    getBodyPart(),
	})
}

func GetEquipment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
		"data":    getEquipment(),
	})
}

func GetTarget(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
		"data":    getTarget(),
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
