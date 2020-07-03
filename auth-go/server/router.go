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
	auth := r.Group("/auth")
	{
		auth.GET("/ping", controller.Ping)
		auth.POST("/register", controller.CreateUser)
		auth.POST("/login", controller.LoginUser)
		auth.GET("/claims", controller.GetUserClaims)
		auth.GET("/swagdocs", controller.SwaggerDocs)
	}

	return r
}
