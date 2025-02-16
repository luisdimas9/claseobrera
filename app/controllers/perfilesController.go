package controllers

import (
	"claseobrera/app/handlers"
	"claseobrera/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetPerfiles(c *fiber.Ctx) error {
	return handlers.HandleGet(c, models.GetPerfiles)
}

func PostPerfiles(c *fiber.Ctx) error {
	return handlers.HandlePost(c, models.PostPerfiles)
}

/*func DeletePerfiles(c *fiber.Ctx) error {
	return handlers.HandleDelete(c, models.DeletePerfilesID)
}*/
