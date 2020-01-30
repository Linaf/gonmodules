package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCarModels(c echo.Context) error {
	dealer := c.Param("dealerName")
	make := c.Param("makeId")

	return c.String(http.StatusOK, "List of model are retrieved for dealer: "+dealer+" and make: "+make)
}
