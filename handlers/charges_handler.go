package handlers

import (
  "net/http"
  "strconv"

  "github.com/AntonioTWize/goledger/repositories"
  "github.com/labstack/echo/v4"
)

type ChargeHandler struct {
  Repository *repositories.ChargeRepository
}

func NewChargeHandler(repo *repositories.ChargeRepository) *ChargeHandler {
  return &ChargeHandler{Repository: repo}
}

func (h *ChargeHandler) CreateCharge(c echo.Context) error {
  var charge repositories.Charge
  if err := c.Bind(&charge); err != nil {
    return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
  }

  if charge.Concept == "" || charge.Amount <= 0 || charge.PaymentMethod == "" || charge.Category == "" || charge.Date == "" {
    return c.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required"})
  }

  if err := h.Repository.CreateCharge(&charge); err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
  }

  return c.JSON(http.StatusCreated, charge)
}

func (h *ChargeHandler) GetAllCharges(c echo.Context) error {
  charges, err := h.Repository.GetAllCharges()
  if err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
  }
  return c.JSON(http.StatusOK, charges)
}

func (h *ChargeHandler) GetChargeByID(c echo.Context) error {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
  }

  charge, err := h.Repository.GetChargeByID(id)
  if err != nil {
    return c.JSON(http.StatusNotFound, map[string]string{"error": "Charge not found"})
  }
  return c.JSON(http.StatusOK, charge)
}

func (h *ChargeHandler) UpdateCharge(c echo.Context) error {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
  }

  var updatedCharge repositories.Charge
  if err := c.Bind(&updatedCharge); err != nil {
    return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
  }

  updatedCharge.ID = id
  if err := h.Repository.UpdateCharge(&updatedCharge); err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
  }

  return c.JSON(http.StatusOK, updatedCharge)
}

func (h *ChargeHandler) DeleteCharge(c echo.Context) error {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
  }

  if err := h.Repository.DeleteCharge(id); err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
  }

  return c.JSON(http.StatusOK, map[string]string{"message": "Charge deleted"})
}
