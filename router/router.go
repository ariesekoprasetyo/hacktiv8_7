package router

import (
	"github.com/ariesekoprasetyo/hacktiv8_7/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	router.POST("/post-order", controllers.OrderPost)
	router.GET("/get-order/:id", controllers.OrderGetDataById)
	router.GET("/get-orders", controllers.OrderGetAllData)
	router.PUT("/update-order/:id", controllers.OrderUpdate)
	router.DELETE("/delete-order/:id", controllers.OrderDelete)
	return
}
