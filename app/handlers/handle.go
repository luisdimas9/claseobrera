package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleGet(c *fiber.Ctx, getDataFunc func() (interface{}, error)) error {
	data, err := getDataFunc()
	if err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}
	return createSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func HandleGetByID(c *fiber.Ctx, getByIDFunc func(int) (interface{}, error)) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return createErrorResponse(c, fiber.StatusBadRequest, "ID inválido, debe ser un número entero", err)
	}
	data, err := getByIDFunc(id)
	if err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}
	return createSuccessResponse(c, fiber.StatusOK, "Datos obtenidos con éxito", data)
}

func HandlePost(c *fiber.Ctx, postDataFunc func(interface{}) error) error {
	var data interface{}
	if err := c.BodyParser(&data); err != nil {
		return createErrorResponse(c, fiber.StatusBadRequest, "Datos de entrada no válidos", err)
	}
	if err := postDataFunc(data); err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}
	return createSuccessResponse(c, fiber.StatusCreated, "La organización se ha creado con éxito", nil)
}

func HandlePut(c *fiber.Ctx, putDataFunc func(interface{}) error) error {
	var data interface{}
	if err := c.BodyParser(&data); err != nil {
		return createErrorResponse(c, fiber.StatusBadRequest, "Datos de entrada no válidos", err)
	}
	if err := putDataFunc(data); err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}
	return createSuccessResponse(c, fiber.StatusCreated, "La organización se ha actualizado con éxito", nil)
}

func HandleDelete(c *fiber.Ctx, deleteByIDFunc func(int) error) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return createErrorResponse(c, fiber.StatusBadRequest, "ID inválido, debe ser un número entero", err)
	}
	if err := deleteByIDFunc(id); err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, "Error al procesar la solicitud", err)
	}
	return createSuccessResponse(c, fiber.StatusCreated, "La organización se ha eliminado con éxito", nil)
}

func HandleLogin(c *fiber.Ctx, loginFunc func(username, password, ipAddress, userAgent string) (string, error)) error {
	var requestData struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
		IPAddress string `json:"ip_address"`
		UserAgent string `json:"user_agent"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return createErrorResponse(c, fiber.StatusBadRequest, "Datos de entrada no válidos", err)
	}

	token, err := loginFunc(requestData.Username, requestData.Password, requestData.IPAddress, requestData.UserAgent)
	if err != nil {
		return createErrorResponse(c, fiber.StatusUnauthorized, "Error al iniciar sesión", err)
	}

	return createSuccessResponse(c, fiber.StatusOK, "Inicio de sesión exitoso", fiber.Map{
		"token": token,
	})
}

func HandleLogout(c *fiber.Ctx, logoutFunc func(token, ipAddress, userAgent string) error) error {
	var requestData struct {
		Token     string `json:"token"`
		IPAddress string `json:"ip_address"`
		UserAgent string `json:"user_agent"`
	}

	if err := c.BodyParser(&requestData); err != nil {
		return createErrorResponse(c, fiber.StatusBadRequest, "Datos de entrada no válidos", err)
	}

	err := logoutFunc(requestData.Token, requestData.IPAddress, requestData.UserAgent)
	if err != nil {
		return createErrorResponse(c, fiber.StatusInternalServerError, "Error al cerrar sesión", err)
	}

	return createSuccessResponse(c, fiber.StatusOK, "Sesión cerrada correctamente", nil)
}

func HandleAuthError(c *fiber.Ctx, err error) error {
	return createErrorResponse(c, fiber.StatusUnauthorized, "Invalid or expired JWT", err)
}

func HandlePermissionError(c *fiber.Ctx, err error) error {
	return createErrorResponse(c, fiber.StatusForbidden, "Insufficient permissions", err)
}

func HandleInternalError(c *fiber.Ctx, err error) error {
	return createErrorResponse(c, fiber.StatusInternalServerError, "Internal server error", err)
}
