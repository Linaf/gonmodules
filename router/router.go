package router

import (
	"github.com/Linaf/gonmodules/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	//e.GET("/lruConfigs", hello)
	api.ApiGroup(e)
	return e
}
