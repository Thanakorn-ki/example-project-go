package main

import (
	"github.com/labstack/echo"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("Authorization") == "551122" {
			return next(c)
		}
		return echo.NewHTTPError(401, "Please provide valid token")
	}
}
