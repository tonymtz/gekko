package server

import (
	"github.com/labstack/echo"
	"github.com/tonymtz/gekko/routes"
)

func router(e *echo.Echo) {
	/*
	 * Home
	 */
	e.Static("/", "static")

	/*
	 * Login
	 */
	e.GET("/login/:provider", routes.Login.Get)
	e.GET("/login/:provider/callback", routes.Login.Callback)
}
