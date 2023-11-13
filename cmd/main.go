package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"io"

	"net/http"
	"net/url"

	"time"

	"Matthieu-OD/card_game_sixty_six/cmd/dbutils"
	"Matthieu-OD/card_game_sixty_six/cmd/game"
	"Matthieu-OD/card_game_sixty_six/cmd/wsutils"

	"github.com/google/uuid"

	_ "modernc.org/sqlite"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// template rendering
type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

// app
func main() {
	sqldb, ctx := dbutils.SetupDB()
	defer sqldb.Close()

	game.GameIDToEventChan = make(map[string]chan string)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Static("/static", "web/assets")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("web/templates/**/*.html")),
	}

	e.Renderer = renderer

	// app routes
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/home", map[string]interface{}{})
	}).Name = "home"

	e.GET("/create-new-game", func(c echo.Context) error {
		return createNewGame(c, sqldb, ctx)
	}).Name = "createNewGame"

	e.GET("/waiting-opponent/:gameid", waitingOpponent).Name = "waitingOpponent"

	e.GET("/join-game/:gameid", joinGame).Name = "joinGame"

	e.GET("/game-full", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/gamefull", map[string]interface{}{})
	}).Name = "gameFull"

	e.GET("/game/:gameid", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/game", map[string]interface{}{})
	}).Name = "game"

	e.GET("/sse/:gameid", sseGame).Name = "gameSSE"

	// TODO: write the JS websocket client using HTMX
	e.GET("/ws/:id", wsGame).Name = "gameWebsocket"

	e.Logger.Fatal(e.Start(":8000"))
}

func createNewGame(c echo.Context, sqldb *sql.DB, ctx context.Context) error {
	gameid := uuid.NewString()

	// create the game in the database
	err := dbutils.CreateEmptyGame(sqldb, ctx, gameid)
	if err != nil {
		c.Logger().Fatal(err)
	}

	waitingOpponentURL := c.Echo().Reverse("waitingOpponent", gameid)
	return c.Redirect(http.StatusPermanentRedirect, waitingOpponentURL)
}

func waitingOpponent(c echo.Context) error {
	gameid := c.Param("gameid")
	return c.Render(http.StatusOK, "views/waiting", map[string]interface{}{
		"JoinGameURL": &url.URL{
			Scheme: c.Scheme(),
			Host:   c.Request().Host,
			Path:   c.Echo().Reverse("joinGame", c.Param("gameid")),
		},
		"SseURL":  c.Echo().Reverse("gameSSE", gameid),
		"GameURL": c.Echo().Reverse("game", gameid),
	})
}

func joinGame(c echo.Context) error {
	gameid := c.Param("gameid")

	eventChan, exists := game.GameIDToEventChan[gameid]
	fmt.Printf("eventChan %v, exists %v", eventChan, exists)

	if exists {
		fmt.Printf("here")
		eventChan <- "start"
		fmt.Printf("here2")
		return c.Redirect(http.StatusPermanentRedirect, c.Echo().Reverse("game", gameid))
	} else {
		fmt.Printf("here3")
		return c.Redirect(http.StatusPermanentRedirect, c.Echo().Reverse("gameFull"))
	}
}

func sseGame(c echo.Context) error {
	gameid := c.Param("gameid")

	eventChan, exists := game.GameIDToEventChan[gameid]

	if !exists {
		eventChan = make(chan string)
		game.GameIDToEventChan[gameid] = eventChan
	}

	fmt.Printf("SSE connection started for game %v", gameid)
	fmt.Printf("game event chan %v", game.GameIDToEventChan)

	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	// TODO: change this for production
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	progress := 0
	for {
		fmt.Printf("progress %v", progress)
		fmt.Fprintf(c.Response(), "message: %v\n", "progress")
		fmt.Fprintf(c.Response(), "data: %v\n\n", progress)
		c.Response().Flush()

		progress++
		time.Sleep(2 * time.Second)
	}

	// for {
	// 	event := <-eventChan
	// 	fmt.Printf("event %v", event)
	// 	fmt.Fprintf(c.Response(), "event: %s\n", event)
	// 	fmt.Fprintf(c.Response(), "data: %s\n", "hello")
	// 	fmt.Fprintf(c.Response(), "id: %s\n\n", gameid)
	// 	c.Response().Flush()
	// }
}

func wsGame(c echo.Context) error {
	conn, err := wsutils.Upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Printf("Websocket upgrade error: %v", err)
		return err
	}
	defer conn.Close()

	for {
		// Write
		err := conn.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, _, err = conn.ReadMessage()
		// _, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)

		}
		// c.Logger().Printf("%s\n", msg)
	}
}
