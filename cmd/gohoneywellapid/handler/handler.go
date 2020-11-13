package handler

import (
	"net/http"

	"github.com/labstack/echo"
	hwapi "github.com/prune998/gohoneywellapi"
)

type Handler struct {
	HwData *hwapi.TSerie
}

func (h *Handler) GetLocation(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, h.HwData)
}

func (h *Handler) GetDevice(c echo.Context) (err error) {
	id := c.Param("id")

	return c.JSON(http.StatusOK, id)
}
