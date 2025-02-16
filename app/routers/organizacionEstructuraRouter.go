package routers

import (
	"claseobrera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func OrganizacionEstructuraRoutes(app *fiber.App) {
	api := app.Group("/api")

	orgGroup := api.Group("/organizacionestructura")

	orgGroup.Get("/", controllers.GetOrganizacionEstructura)
	orgGroup.Get("/:id", controllers.GetOrganizacionEstructuraID)
	orgGroup.Post("/", controllers.PostOrganizacionEstructura)
	orgGroup.Put("/", controllers.PutOrganizacionEstructura)
	orgGroup.Delete("/:id", controllers.DeleteOrganizacionEstructura)
}
