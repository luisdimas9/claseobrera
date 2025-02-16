package controllers

import (
	"claseobrera/app/handlers"
	"claseobrera/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetParroquia(c *fiber.Ctx) error {
	return handlers.HandleGet(c, models.GetParroquia)
}

func GetParroquiaID(c *fiber.Ctx) error {
	return handlers.HandleGetByID(c, models.GetParroquiaID)
}

func PostParroquia(c *fiber.Ctx) error {
	return handlers.HandlePost(c, models.PostParroquia)
}

func PutParroquia(c *fiber.Ctx) error {
	return handlers.HandlePut(c, models.PutParroquia)
}

func DeleteParroquia(c *fiber.Ctx) error {
	return handlers.HandleDelete(c, models.DeleteParroquia)
}
