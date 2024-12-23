package handlers

import (
    "net/http"

    "github.com/labstack/echo/v4"
)

type UserHandler struct {
    // Puedes agregar dependencias aquí (por ejemplo, servicios o repositorios).
}

func NewUserHandler() *UserHandler {
    return &UserHandler{}
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
    id := c.Param("id")
    return c.JSON(http.StatusOK, map[string]string{
        "id":      id,
        "message": "User details retrieved",
    })
}
