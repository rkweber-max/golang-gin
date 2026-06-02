package controllers

import "github.com/gin-gonic/gin"

func FindAllClassmates(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Rodrigo",
	})
}

func Hello(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"API say": "Hello " + name,
	})
}
