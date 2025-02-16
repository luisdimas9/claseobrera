package routers

import (
	"claseobrera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func OrganizacionGeneralRoutes(app *fiber.App) {
	api := app.Group("/api")

	orgGroup := api.Group("/organizaciongeneral")

	orgGroup.Get("/", controllers.GetOrganizacionGeneral)
	orgGroup.Get("/:id", controllers.GetOrganizacionGeneralID)
	orgGroup.Post("/", controllers.PostOrganizacionGeneral)
	orgGroup.Put("/", controllers.PutOrganizacionGeneral)
	orgGroup.Delete("/:id", controllers.DeleteOrganizacionGeneral)
}
