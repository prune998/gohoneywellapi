package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	hwapi "github.com/prune998/gohoneywellapi"
)

type Handler struct {
	Hw *hwapi.Honeywellapi
}

func (h *Handler) GetLocations(c echo.Context) (err error) {
	locations, err := h.Hw.GetLocations()
	if err != nil {
		c.Logger().Errorf("error getting Locations: %s", err)

	}
	return c.JSON(http.StatusOK, locations)
}

func (h *Handler) GetLocation(c echo.Context) (err error) {
	locationid := c.Param("locationid")

	locations, err := h.Hw.GetLocations()
	if err != nil {
		c.Logger().Errorf("error getting Locations: %s", err)
	}

	for _, loc := range locations {
		if strconv.Itoa(loc.LocationID) == locationid {
			return c.JSON(http.StatusOK, loc)
		}
	}
	return c.JSON(http.StatusNotFound, nil)
}

func (h *Handler) GetDevices(c echo.Context) (err error) {
	locationID := c.Param("locationid")

	locations, err := h.Hw.GetLocations()
	if err != nil {
		c.Logger().Errorf("error getting Locations: %s", err)
	}

	for _, loc := range locations {
		if strconv.Itoa(loc.LocationID) == locationID {
			return c.JSON(http.StatusOK, loc.Devices)
		}
	}
	return c.JSON(http.StatusNotFound, nil)
}

func (h *Handler) GetDevice(c echo.Context) (err error) {
	locationID := c.Param("locationid")
	deviceID := c.Param("deviceid")

	locations, err := h.Hw.GetLocations()
	if err != nil {
		c.Logger().Errorf("error getting Locations: %s", err)
	}

	for _, loc := range locations {
		if strconv.Itoa(loc.LocationID) == locationID {

			for _, dev := range loc.Devices {
				if dev.DeviceID == deviceID {
					return c.JSON(http.StatusOK, dev)
				}
			}
		}
	}
	return c.JSON(http.StatusNotFound, nil)
}
