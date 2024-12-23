package handlers

import (
  "net/http"

  "github.com/labstack/echo/v4"
)

type Charge struct {
  ID           string  `json:"id"`
  Concept      string  `json:"concept"`
  Amount       float64 `json:"amount"`
  PaymentMethod string  `json:"payment_method"`
  Category     string  `json:"category"`
  Date         string  `json:"date"`
}

type ChargeHandler struct {
  // Aquí puedes agregar dependencias como servicios o repositorios en el futuro.
}

func NewChargeHandler() *ChargeHandler {
  return &ChargeHandler{}
}

func (h *ChargeHandler) CreateCharge(c echo.Context) error {
  charge := new(Charge)
  if err := c.Bind(charge); err != nil {
    return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
  }
  charge.ID = "1" // Generar ID único en el futuro
  return c.JSON(http.StatusCreated, charge)
}

func (h *ChargeHandler) GetAllCharges(c echo.Context) error {
  // Datos de ejemplo
  charges := []Charge{
    {ID: "1", Concept: "Compra", Amount: 100.0, PaymentMethod: "Tarjeta", Category: "Compras", Date: "2024-12-22"},
    {ID: "2", Concept: "Transporte", Amount: 50.0, PaymentMethod: "Efectivo", Category: "Viajes", Date: "2024-12-21"},
  }
  return c.JSON(http.StatusOK, charges)
}

func (h *ChargeHandler) GetChargeByID(c echo.Context) error {
  id := c.Param("id")
  // Datos de ejemplo
  if id == "1" {
    charge := Charge{ID: "1", Concept: "Compra", Amount: 100.0, PaymentMethod: "Tarjeta", Category: "Compras", Date: "2024-12-22"}
    return c.JSON(http.StatusOK, charge)
  }
  return c.JSON(http.StatusNotFound, map[string]string{"error": "Charge not found"})
}

func (h *ChargeHandler) UpdateCharge(c echo.Context) error {
  id := c.Param("id")
  updatedCharge := new(Charge)
  if err := c.Bind(updatedCharge); err != nil {
    return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
  }
  updatedCharge.ID = id // Mantener el mismo ID
  return c.JSON(http.StatusOK, updatedCharge)
}

func (h *ChargeHandler) DeleteCharge(c echo.Context) error {
  id := c.Param("id")
  return c.JSON(http.StatusOK, map[string]string{"message": "Charge deleted", "id": id})
}
