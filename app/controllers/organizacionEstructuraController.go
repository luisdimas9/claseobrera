package controllers

import (
	"claseobrera/app/handlers"
	"claseobrera/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetOrganizacionEstructura(c *fiber.Ctx) error {
	return handlers.HandleGet(c, models.GetOrganizacionEstructura)
}

func GetOrganizacionEstructuraID(c *fiber.Ctx) error {
	return handlers.HandleGetByID(c, models.GetOrganizacionEstructuraID)
}

func PostOrganizacionEstructura(c *fiber.Ctx) error {
	return handlers.HandlePost(c, models.PostOrganizacionEstructura)
}

func PutOrganizacionEstructura(c *fiber.Ctx) error {
	return handlers.HandlePut(c, models.PutOrganizacionEstructura)
}

func DeleteOrganizacionEstructura(c *fiber.Ctx) error {
	return handlers.HandleDelete(c, models.DeleteOrganizacionEstructuraID)
}
