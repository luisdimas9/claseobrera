package routers

import (
	"claseobrera/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func SessionRoutes(app *fiber.App) {
	api := app.Group("/api")

	sessionGroup := api.Group("/session")

	sessionGroup.Post("/login", controllers.PostLogin)
	sessionGroup.Post("/logout", controllers.PostLogout)
}
