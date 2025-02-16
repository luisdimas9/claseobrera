package routers

import (
	"claseobrera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PermisosRoutes(app *fiber.App) {
	api := app.Group("/api")

	permisosGroup := api.Group("/permisos")

	permisosGroup.Get("/", controllers.GetPermisos)
	permisosGroup.Get("/:id", controllers.GetPermisosID)
	permisosGroup.Post("/", controllers.PostPermisos)
	permisosGroup.Put("/", controllers.PutPermisos)
	permisosGroup.Delete("/:id", controllers.DeletePermisos)
}
