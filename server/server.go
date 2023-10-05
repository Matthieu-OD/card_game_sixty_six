package main

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Static("/static", "assets")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/**/*.html")),
	}

	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/home", map[string]interface{}{})
	})
	e.GET("/waiting-opponent", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/waiting", map[string]interface{}{})
	})
	e.GET("/game", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/game", map[string]interface{}{})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
