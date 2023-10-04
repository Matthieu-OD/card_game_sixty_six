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

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}

	e.Renderer = renderer

	// viteURL, _ := url.Parse("http://localhost:5173")
	// proxy := httputil.NewSingleHostReverseProxy(viteURL)
	//
	// e.GET("/", func(c echo.Context) error {
	// 	proxy.ServeHTTP(c.Response(), c.Request())
	// 	return nil
	// })

	e.GET("/something", func(c echo.Context) error {
		return c.Render(http.StatusOK, "template.html", map[string]interface{}{
			"name": "Dolly!",
		})
	}).Name = "foobar"

	e.Logger.Fatal(e.Start(":8000"))
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}
