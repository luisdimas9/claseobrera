package controllers

import (
	"claseobrera/app/handlers"
	"claseobrera/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetPermisos(c *fiber.Ctx) error {
	return handlers.HandleGet(c, models.GetPermisos)
}

func GetPermisosID(c *fiber.Ctx) error {
	return handlers.HandleGetByID(c, models.GetPermisosID)
}

func PostPermisos(c *fiber.Ctx) error {
	return handlers.HandlePost(c, models.PostPermisos)
}

func PutPermisos(c *fiber.Ctx) error {
	return handlers.HandlePut(c, models.PutPermisos)
}

func DeletePermisos(c *fiber.Ctx) error {
	return handlers.HandleDelete(c, models.DeletePermisosID)
}
