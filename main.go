package main

import (
	"github.com/gofiber/fiber/v2"
	"scylladb/controllers"
	"scylladb/db"
	"scylladb/repositories"
	"scylladb/routers"
	"scylladb/services"
)

func main() {
	db.Init()

	app := fiber.New()

	todoRepository := &repositories.TODORepository{}
	todoService := &services.TODOService{Repository: todoRepository}
	todoController := &controllers.TODOController{Service: todoService}

	routes.SetupTODORoutes(app, todoController)

	app.Listen(":3000")
}
