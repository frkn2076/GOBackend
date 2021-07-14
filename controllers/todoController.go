package controllers

import (
	"app/GOBackend/controllers/request"
	"app/GOBackend/database"

	"github.com/gin-gonic/gin"
)

func AddItem(context *gin.Context) {
	var todoRequest request.TodoRequest
	if err := context.ShouldBindJSON(&todoRequest); err != nil {
		context.JSON(200, gin.H{
			"IsError": true,
			"ErrorMessage": "Geçici süre hizmet veremiyoruz",
		})
		return
	}

	todo := database.Todo{Name: todoRequest.Name}

	database.CreateTodo(todo)

	context.JSON(200, gin.H{
		"IsError": false,
		"ErrorMessage": nil,
	})
}

func Items(context *gin.Context) {
	todos := database.GetAllTodos()

	context.JSON(200, gin.H{
		"IsError": false,
		"Items": todos,
	})
}
