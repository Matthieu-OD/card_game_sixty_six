package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	// "Matthieu-OD/card_game_sixty_six/server/redis"

	"github.com/google/uuid"
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

	// list all the routes of the application
	e.GET("/", func(c echo.Context) error {
		createNewGameURL := c.Echo().Reverse("createNewGame")
		return c.Render(http.StatusOK, "views/home", map[string]interface{}{
			"CreateNewGameURL": createNewGameURL,
		})
	}).Name = "home"
	e.GET("/create-new-game", createNewGame).Name = "createNewGame"
	e.GET("/waiting-opponent/:gameid", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/waiting", map[string]interface{}{})
	}).Name = "waitingOpponent"
	e.GET("/game", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/game", map[string]interface{}{})
	}).Name = "game"
	e.GET("/ws/:id", wsGame).Name = "gameWebsocket"

	e.Logger.Fatal(e.Start(":8000"))
}

var (
	upgrader = websocket.Upgrader{}
)

func createNewGame(c echo.Context) error {
	// TODO: create a new uuid and add it to the redis server

	gameID := uuid.NewString()
	waitingOpponentURL := c.Echo().Reverse("waitingOpponent", gameID)
	return c.Redirect(http.StatusPermanentRedirect, waitingOpponentURL)
}

func wsGame(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

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
		log.Printf("%s\n", msg)
	}
}
