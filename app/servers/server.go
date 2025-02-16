package servers

import (
	"claseobrera/app/routers"

	"github.com/gofiber/fiber/v2"
)

func SetupAndListen() {
	app := fiber.New()

	routers.OrganizacionGeneralRoutes(app)
	routers.OrganizacionJerarquiaRoutes(app)
	routers.OrganizacionEstructuraRoutes(app)
	routers.EstadosRoutes(app)
	routers.MunicipiosRoutes(app)
	routers.ParroquiasRoutes(app)
	routers.CiudadesRoutes(app)
	routers.InstitucionesRoutes(app)
	routers.SessionRoutes(app)
	routers.UsersRoutes(app)
	routers.PermisosRoutes(app)
	routers.RolesRoutes(app)
	routers.PerfilesRoutes(app)

	app.Listen(":3000")
}
