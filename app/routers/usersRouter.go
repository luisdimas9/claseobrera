package routers

import (
	"claseobrera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func UsersRoutes(app *fiber.App) {
	api := app.Group("/api")

	usersGroup := api.Group("/users")

	usersGroup.Get("/", controllers.GetUsers)
	usersGroup.Get("/:id", controllers.GetUsersID)
	usersGroup.Post("/", controllers.PostUsers)
	usersGroup.Put("/", controllers.PutUsers)
	usersGroup.Delete("/:id", controllers.DeleteUsers)
}
