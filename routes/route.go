package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rkweber-max/golang-gin/controllers"
)

func HandleFunc() {
	r := gin.Default()
	r.GET("/classmates", controllers.FindAllClassmates)

	r.Run()
}
