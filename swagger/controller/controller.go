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
		"http://172.28.1.1:8000/auth/swagdocs",
	}, {
		2,
		"Fetch",
		"http://172.28.1.1:8000/swagger/swagger.json",
	}}

	return c.Render(http.StatusOK, "index.html", swaggers)
}
