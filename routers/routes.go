package routes

import (
	"github.com/gofiber/fiber/v2"
	"scylladb/controllers"
)

func SetupTODORoutes(app *fiber.App, controller *controllers.TODOController) {
	app.Post("/todos", controller.CreateTODO)
	app.Get("/todos/:user_id/:id", controller.GetTODO)
	app.Get("/todos/:user_id", controller.ListTODOs)
	app.Put("/todos/:user_id/:id", controller.UpdateTODO)
	app.Delete("/todos/:user_id/:id", controller.DeleteTODO)
}
