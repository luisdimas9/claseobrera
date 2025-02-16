package routers

import (
	"claseobrera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func ParroquiasRoutes(app *fiber.App) {
	api := app.Group("/api")

	parroquiaGroup := api.Group("/parroquia")

	parroquiaGroup.Get("/", controllers.GetParroquia)
	parroquiaGroup.Get("/:id", controllers.GetParroquiaID)
	parroquiaGroup.Post("/", controllers.PostParroquia)
	parroquiaGroup.Put("/", controllers.PutParroquia)
	parroquiaGroup.Delete("/:id", controllers.DeleteParroquia)
}
