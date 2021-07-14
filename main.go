package main

import(
	"app/GOBackend/database"
	"app/GOBackend/router"
)

func main() {
	// Dummy
	todo := database.Todo{Name: "Read book"}
	database.CreateTodo(todo)

	r := router.SetupRouter()
	r.Run(":5000")
}