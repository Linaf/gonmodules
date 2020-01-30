package api

import (
	"github.com/Linaf/gonmodules/api/handlers"
	"github.com/labstack/echo/v4"
)

func ApiGroup(e *echo.Echo) {
	e.GET("/dealer/:dealerName/make/:makeId/models", handlers.GetCarModels)
}
