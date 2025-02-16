package controllers

import (
	"claseobrera/app/handlers"
	"claseobrera/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	return handlers.HandleGet(c, models.GetUsers)
}

func GetUsersID(c *fiber.Ctx) error {
	return handlers.HandleGetByID(c, models.GetUsersID)
}

func PostUsers(c *fiber.Ctx) error {
	return handlers.HandlePost(c, models.PostUsers)
}

func PutUsers(c *fiber.Ctx) error {
	return handlers.HandlePut(c, models.PutUsers)
}

func DeleteUsers(c *fiber.Ctx) error {
	return handlers.HandleDelete(c, models.DeleteUsersID)
}
