package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// H -
type H map[string]interface{}

// OK wraps ok request response
func OK(c echo.Context, msg string) error {
	return c.JSON(
		http.StatusOK,
		H{
			"success": true,
			"status":  http.StatusOK,
			"msg":     msg,
		},
	)
}

// Data wraps request response with data
func Data(c echo.Context, data interface{}) error {
	return c.JSON(
		http.StatusOK,
		H{
			"success": true,
			"status":  http.StatusOK,
			"msg":     "OK",
			"data":    data,
		},
	)
}

// Bad wraps bad request response
func Bad(c echo.Context) error {
	return echo.NewHTTPError(
		http.StatusBadRequest,
		H{
			"success": false,
			"status":  http.StatusBadRequest,
			"msg":     "Bad Request",
		},
	)
}

// Unprocessable wraps unprocessable entity response
func Unprocessable(c echo.Context, msg string) error {
	return echo.NewHTTPError(
		http.StatusUnprocessableEntity,
		H{
			"success": false,
			"status":  http.StatusUnprocessableEntity,
			"msg":     msg,
		},
	)
}

// InternalError wraps internal error response
func InternalError(c echo.Context) error {
	return echo.NewHTTPError(
		http.StatusInternalServerError,
		H{
			"success": false,
			"status":  http.StatusInternalServerError,
			"msg":     "Internal Server Error",
		},
	)
}

// NotFound wraps not found request response
func NotFound(c echo.Context) error {
	return echo.NewHTTPError(
		http.StatusNotFound,
		H{
			"success": false,
			"status":  http.StatusNotFound,
			"msg":     "Not found",
		},
	)
}

// Forbidden wraps forbidden request response
func Forbidden(c echo.Context) error {
	return echo.NewHTTPError(
		http.StatusForbidden,
		H{
			"success": false,
			"status":  http.StatusForbidden,
			"msg":     "Forbidden",
		},
	)
}

// Unauthorized wraps unauthorized request response
func Unauthorized(c echo.Context) error {
	return echo.NewHTTPError(
		http.StatusUnauthorized,
		H{
			"success": false,
			"status":  http.StatusUnauthorized,
			"msg":     "Unauthorized",
		},
	)
}
