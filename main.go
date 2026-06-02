package main

import "github.com/gin-gonic/gin"

func FindAllClassmates(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Rodrigo",
	})
}

func main() {
	r := gin.Default()
	r.GET("/classmates", FindAllClassmates)

	r.Run()
}
