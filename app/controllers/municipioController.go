package controllers

import (
	"claseobrera/app/handlers"
	"claseobrera/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetMunicipio(c *fiber.Ctx) error {
	return handlers.HandleGet(c, models.GetMunicipio)
}

func GetMunicipioID(c *fiber.Ctx) error {
	return handlers.HandleGetByID(c, models.GetMunicipioID)
}

func PostMunicipio(c *fiber.Ctx) error {
	return handlers.HandlePost(c, models.PostMunicipio)
}

func PutMunicipio(c *fiber.Ctx) error {
	return handlers.HandlePut(c, models.PutMunicipio)
}

func DeleteMunicipio(c *fiber.Ctx) error {
	return handlers.HandleDelete(c, models.DeleteMunicipio)
}
