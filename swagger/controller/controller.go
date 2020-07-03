package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Swagger struct {
	ID   int    `json:"_id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Url  string `json:"url" bson:"url"`
}

func Index(c echo.Context) error {
	swaggers := []Swagger{{
		1,
		"Auth",
		"http://localhost:7000/auth/swagdocs",
	}, {
		2,
		"Fetch",
		"http://localhost:8000/fetch/swagdocs",
	}}

	return c.Render(http.StatusOK, "index.html", swaggers)
}
