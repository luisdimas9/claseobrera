package controllers

import (
	"claseobrera/app/handlers"
	"claseobrera/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetOrganizacionGeneral(c *fiber.Ctx) error {
	return handlers.HandleGet(c, models.GetOrganizacionGeneral)
}

func GetOrganizacionGeneralID(c *fiber.Ctx) error {
	return handlers.HandleGetByID(c, models.GetOrganizacionGeneralID)
}

func PostOrganizacionGeneral(c *fiber.Ctx) error {
	return handlers.HandlePost(c, models.PostOrganizacionGeneral)
}

func PutOrganizacionGeneral(c *fiber.Ctx) error {
	return handlers.HandlePut(c, models.PutOrganizacionGeneral)
}

func DeleteOrganizacionGeneral(c *fiber.Ctx) error {
	return handlers.HandleDelete(c, models.DeleteOrganizacionGeneralID)
}
