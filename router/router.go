package router

import (
	"github.com/ariesekoprasetyo/hacktiv8_7/controllers"
	"github.com/ariesekoprasetyo/hacktiv8_7/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() (router *gin.Engine) {
	router = gin.Default()
	v1route := router.Group("/api/v1")
	v1route.Use(
		middleware.CORSMiddleware,
	)
	{
		order := v1route.Group("/order")
		{
			order.POST("/post-order", controllers.OrderPost)
			order.GET("/get-order/:id", controllers.OrderGetDataById)
			order.GET("/get-orders", controllers.OrderGetAllData)
			order.PUT("/update-order/:id", controllers.OrderUpdate)
			order.DELETE("/delete-order/:id", controllers.OrderDelete)
		}
	}

	return
}
