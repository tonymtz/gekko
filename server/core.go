package server

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/labstack/echo/engine/fasthttp"

	"github.com/tonymtz/gekko/server/config"
)

func Start() {
	log.Print("Starting a new gekko service")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router(e)

	port := strconv.Itoa(config.Config.Port)

	log.Print("Serving on port " + port)
	e.Run(fasthttp.New(":" + port))
}
