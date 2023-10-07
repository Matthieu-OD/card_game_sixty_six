package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"Matthieu-OD/card_game_sixty_six/server/redis"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "assets")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/**/*.html")),
	}

	e.Renderer = renderer

	// TODO: move in the game view
	ctx := context.Background()

	err := client.HSet(ctx, "game:123")
	// TODO: replace this with REDIS
	var (
		userNumber = 0
	)

	// list all the routes of the application
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/home", map[string]interface{}{})
	})
	e.GET("/waiting-opponent", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/waiting", map[string]interface{}{})
	})
	e.GET("/game", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/game", map[string]interface{}{})
	})
	e.GET("/ws/:id", wsGame)

	e.Logger.Fatal(e.Start(":8000"))
}

var (
	upgrader = websocket.Upgrader{}
)

func wsGame(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	errId := ws.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(userNumber)))
	if errId != nil {
		c.Logger().Error(err)
	}
	userNumber += 1

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)

		}
		fmt.Printf("%s\n", msg)
	}
}
