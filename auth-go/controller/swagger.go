package controller

import (
	"github.com/labstack/echo"
)

// @Tags Swagger
// @Summary get Swagger Documentation
// @Router /auth/swagdocs [get]
func SwaggerDocs(c echo.Context) error {
	c.Request().Header.Set("Content-Description", "File Transfer")
	c.Request().Header.Set("Content-Transfer-Encoding", "binary")
	c.Request().Header.Set("Content-Disposition", "attachment; filename=swagger.json")
	c.Request().Header.Set("Content-Type", "application/octet-stream")
	return c.File("docs/swagger.json")
}
