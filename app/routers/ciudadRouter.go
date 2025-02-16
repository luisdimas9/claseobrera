package routers

import (
	"claseobrera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func CiudadesRoutes(app *fiber.App) {
	api := app.Group("/api")

	ciudadGroup := api.Group("/ciudad")

	ciudadGroup.Get("/", controllers.GetCiudad)
	ciudadGroup.Get("/:id", controllers.GetCiudadID)
	ciudadGroup.Post("/", controllers.PostCiudad)
	ciudadGroup.Put("/", controllers.PutCiudad)
	ciudadGroup.Delete("/:id", controllers.DeleteCiudad)
}
