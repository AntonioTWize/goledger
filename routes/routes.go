package routes

import (
  "github.com/labstack/echo/v4"
  "github.com/AntonioTWize/goledger/handlers"
)

func NewRouter(chargeHandler *handlers.ChargeHandler) *echo.Echo {
  e := echo.New()

  // Grupo de rutas para "charges"
  charges := e.Group("/charges")
  charges.POST("", chargeHandler.CreateCharge)       // Crear un cargo
  charges.GET("", chargeHandler.GetAllCharges)       // Obtener todos los cargos
  charges.GET("/:id", chargeHandler.GetChargeByID)   // Obtener un cargo por ID
  charges.PUT("/:id", chargeHandler.UpdateCharge)    // Actualizar un cargo por ID
  charges.DELETE("/:id", chargeHandler.DeleteCharge) // Eliminar un cargo por ID

  return e
}
