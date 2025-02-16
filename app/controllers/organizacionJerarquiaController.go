package controllers

import (
	"claseobrera/app/handlers"
	"claseobrera/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetOrganizacionJerarquia(c *fiber.Ctx) error {
	return handlers.HandleGet(c, models.GetOrganizacionJerarquia)
}

func GetOrganizacionJerarquiaID(c *fiber.Ctx) error {
	return handlers.HandleGetByID(c, models.GetOrganizacionJerarquiaID)
}

func PostOrganizacionJerarquia(c *fiber.Ctx) error {
	return handlers.HandlePost(c, models.PostOrganizacionJerarquia)
}

func PutOrganizacionJerarquia(c *fiber.Ctx) error {
	return handlers.HandlePut(c, models.PutOrganizacionJerarquia)
}

func DeleteOrganizacionJerarquia(c *fiber.Ctx) error {
	return handlers.HandleDelete(c, models.DeleteOrganizacionJerarquiaID)
}
