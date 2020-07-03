package server

import (
	"html/template"
	"io"

	"github.com/radhianamri/efishery-project/swagger/controller"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func MainRoutes() *echo.Echo {

	e := echo.New()
	e.Debug = true
	e.Renderer = &TemplateRenderer{
		templates: template.Must(template.ParseGlob("resource/views/*.html")),
	}
	e.Static("/assets", "resource/assets")
	e.GET("/", controller.Index)
	return e
}
