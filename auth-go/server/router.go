package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/radhianamri/efishery-project/auth-go/config"
	"github.com/radhianamri/efishery-project/auth-go/controller"
)

// NewRouter returns *echo.Echo after setting up the routes
func NewRouter() *echo.Echo {
	r := echo.New()
	conf := config.GetConfig()

	if conf.Mode != "prod" {
		r.Debug = true
	}

	r.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     conf.CorsURL,
			AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions},
			AllowCredentials: true,
		}),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, latency=${latency}, status=${status}\n",
		}),
	)
	// Ping server
	r.GET("/ping", controller.Ping)

	return r
}
