package main

import (
	"errors"

	"github.com/labstack/echo"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header["Authorization"][0] == "551122" {
			return next(c)
		} else {
			return errors.New("test")
		}

	}
}