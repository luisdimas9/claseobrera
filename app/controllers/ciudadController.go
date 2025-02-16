package controllers

import (
	"claseobrera/app/handlers"
	"claseobrera/app/models"
	"claseobrera/security/interfaz"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetCiudad(c *fiber.Ctx) error {
	return handlers.HandleGet(c, func() (interface{}, error) {
		return models.GetCiudad()
	})
}

func GetCiudadID(c *fiber.Ctx) error {

	idStr := c.Params("id")
	sanitizedID, err := interfaz.AuthorizeData(idStr)
	if err != nil {
		return handlers.HandleInternalError(c, errors.New(err.Error()))
	}

	intID, err := strconv.Atoi(sanitizedID)
	if err != nil {
		return handlers.HandleInternalError(c, errors.New("ID inválido, debe ser un número entero"))
	}

	return handlers.HandleGetByID(c, func(_ int) (interface{}, error) {
		return models.GetCiudadID(intID)
	})
}

func PostCiudad(c *fiber.Ctx) error {
	return handlers.HandlePost(c, func(data interface{}) error {
		sanitizedData, err := interfaz.AuthorizeData(data)
		if err != nil {
			return handlers.HandleInternalError(c, errors.New(err.Error()))
		}
		return models.PostCiudad(sanitizedData)
	})
}
func PutCiudad(c *fiber.Ctx) error {
	return handlers.HandlePut(c, func(data interface{}) error {
		sanitizedData, err := interfaz.AuthorizeData(data)
		if err != nil {
			return handlers.HandleInternalError(c, errors.New(err.Error()))
		}
		return models.PutCiudad(sanitizedData)
	})
}

func DeleteCiudad(c *fiber.Ctx) error {
	return handlers.HandleDelete(c, models.DeleteCiudad)
}
