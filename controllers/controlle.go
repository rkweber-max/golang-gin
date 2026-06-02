package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rkweber-max/golang-gin/models"
)

func FindAllClassmates(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Hello(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API say": "Hello " + name,
	})
}
