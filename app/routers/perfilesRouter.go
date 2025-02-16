package routers

import (
	"claseobrera/app/controllers"
	"claseobrera/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func PerfilesRoutes(app *fiber.App) {
	api := app.Group("/api")
	corsOptions := middleware.CorsOptions{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"X-Custom-Header"},
		MaxAge:           3600,
	}

	// Se aplica el middleware de CORS, sin que el router conozca nada sobre el módulo security.
	api.Use(middleware.CorsMiddleware(corsOptions))
	api.Use(middleware.AuthRequired)
	api.Use(middleware.CheckPermissions)

	perfilesGroup := api.Group("/perfiles")

	perfilesGroup.Get("/", controllers.GetPerfiles)
	perfilesGroup.Post("/", controllers.PostPerfiles)
}
