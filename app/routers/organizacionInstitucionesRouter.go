package routers

import (
	"claseobrera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func InstitucionesRoutes(app *fiber.App) {
	api := app.Group("/api")

	institucionesGroup := api.Group("/instituciones")

	institucionesGroup.Get("/", controllers.GetInstituciones)
	institucionesGroup.Get("/:id", controllers.GetInstitucionesID)
	institucionesGroup.Post("/", controllers.PostInstituciones)
	institucionesGroup.Put("/", controllers.PutInstituciones)
	institucionesGroup.Delete("/:id", controllers.DeleteInstituciones)
}
