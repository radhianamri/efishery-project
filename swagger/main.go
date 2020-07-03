package main

import (
	"github.com/radhianamri/efishery-project/swagger/server"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// config.Init()
	e := server.MainRoutes()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Start(":9000")
}
