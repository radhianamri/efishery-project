package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type H map[string]interface{}

// Ping returns pong
func Ping(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		H{
			"success": true,
			"status":  http.StatusOK,
			"msg":     "pong",
		},
	)
}
