package router

import (
	"app/GOBackend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	todoRoute := router.Group("/todo")
	{
		todoRoute.POST("add", controllers.AddItem)
		todoRoute.GET("items", controllers.Items)
	}

	return router
}
