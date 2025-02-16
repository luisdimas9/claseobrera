package routers

import (
	"claseobrera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func EstadosRoutes(app *fiber.App) {
	api := app.Group("/api")

	estadoGroup := api.Group("/estado")

	estadoGroup.Get("/", controllers.GetEstado)
	estadoGroup.Get("/:id", controllers.GetEstadoID)
	estadoGroup.Post("/", controllers.PostEstado)
	estadoGroup.Put("/", controllers.PutEstado)
	estadoGroup.Delete("/:id", controllers.DeleteEstado)
}
