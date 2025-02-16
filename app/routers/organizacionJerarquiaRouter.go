package routers

import (
	"claseobrera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func OrganizacionJerarquiaRoutes(app *fiber.App) {
	api := app.Group("/api")

	orgGroup := api.Group("/organizacionjerarquia")

	orgGroup.Get("/", controllers.GetOrganizacionJerarquia)
	orgGroup.Get("/:id", controllers.GetOrganizacionJerarquiaID)
	orgGroup.Post("/", controllers.PostOrganizacionJerarquia)
	orgGroup.Put("/", controllers.PutOrganizacionJerarquia)
	orgGroup.Delete("/:id", controllers.DeleteOrganizacionJerarquia)
}
