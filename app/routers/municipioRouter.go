package routers

import (
	"claseobrera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func MunicipiosRoutes(app *fiber.App) {
	api := app.Group("/api")

	municipioGroup := api.Group("/municipio")

	municipioGroup.Get("/", controllers.GetMunicipio)
	municipioGroup.Get("/:id", controllers.GetMunicipioID)
	municipioGroup.Post("/", controllers.PostMunicipio)
	municipioGroup.Put("/", controllers.PutMunicipio)
	municipioGroup.Delete("/:id", controllers.DeleteMunicipio)
}
