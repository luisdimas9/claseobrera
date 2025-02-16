package controllers

import (
	"claseobrera/app/handlers"
	"claseobrera/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetRoles(c *fiber.Ctx) error {
	return handlers.HandleGet(c, models.GetRoles)
}

func GetRolesID(c *fiber.Ctx) error {
	return handlers.HandleGetByID(c, models.GetRolesID)
}

func PostRoles(c *fiber.Ctx) error {
	return handlers.HandlePost(c, models.PostRoles)
}

func PutRoles(c *fiber.Ctx) error {
	return handlers.HandlePut(c, models.PutRoles)
}

func DeleteRoles(c *fiber.Ctx) error {
	return handlers.HandleDelete(c, models.DeleteRolesID)
}
