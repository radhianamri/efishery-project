package controller

import (
	"github.com/labstack/echo"
)

// Ping returns pong
func Ping(c echo.Context) error {
	return OK(c, "pong")
}
