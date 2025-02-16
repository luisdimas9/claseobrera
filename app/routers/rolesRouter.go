package routers

import (
	"claseobrera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func RolesRoutes(app *fiber.App) {
	api := app.Group("/api")

	rolesGroup := api.Group("/roles")

	rolesGroup.Get("/", controllers.GetRoles)
	rolesGroup.Get("/:id", controllers.GetRolesID)
	rolesGroup.Post("/", controllers.PostRoles)
	rolesGroup.Put("/", controllers.PutRoles)
	rolesGroup.Delete("/:id", controllers.DeleteRoles)
}
