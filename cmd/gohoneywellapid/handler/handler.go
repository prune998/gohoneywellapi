package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	hwapi "github.com/prune998/gohoneywellapi"
)

type Handler struct {
	Hw *hwapi.HoneywellAPI
}

//GetLocations get all the data from locations
func (h *Handler) GetLocations(c echo.Context) (err error) {
	locations, err := h.Hw.GetLocations()
	if err != nil {
		c.Logger().Errorf("error getting Locations: %s", err)

	}
	return c.JSON(http.StatusOK, locations)
}

// GetLocation get specific location data
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

// GetDevices get all the devices for a location
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

// GetDevice get data of a specific device
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

// GetSchedule get the scedule applied on a specific device
func (h *Handler) GetSchedule(c echo.Context) (err error) {
	locationID := c.Param("locationid")
	deviceID := c.Param("deviceid")

	s, err := h.Hw.GetSchedule(locationID, deviceID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, s)
}
