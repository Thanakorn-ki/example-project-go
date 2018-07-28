package main

import (
	"github.com/labstack/echo"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header["Authorization"][0] == "551122" {
			return next(c)
		} else {
			return echo.NewHTTPError(401, "Please provide valid token")
		}
	}
}
