package main

import(
	_ "app/GOBackend/database"
	"app/GOBackend/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":5000")
}