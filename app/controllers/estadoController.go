package controllers

import (
	"claseobrera/app/handlers"
	"claseobrera/app/models"

	"github.com/gofiber/fiber/v2"
)

func GetEstado(c *fiber.Ctx) error {
	return handlers.HandleGet(c, models.GetEstados)
}

func GetEstadoID(c *fiber.Ctx) error {
	return handlers.HandleGetByID(c, models.GetEstadoID)
}

func PostEstado(c *fiber.Ctx) error {
	return handlers.HandlePost(c, models.PostEstado)
}

func PutEstado(c *fiber.Ctx) error {
	return handlers.HandlePut(c, models.PutEstado)
}

func DeleteEstado(c *fiber.Ctx) error {
	return handlers.HandleDelete(c, models.DeleteEstadoID)
}
