package server

import "github.com/labstack/echo"

func router(e *echo.Echo) {
	/*
	 * Home
	 */
	e.Static("/", "static")
}
