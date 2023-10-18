package main

import (
	"context"
	"database/sql"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"

	"Matthieu-OD/card_game_sixty_six/cmd/sql"

	"github.com/google/uuid"

	_ "modernc.org/sqlite"

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
	sqldb, ctx := db.SetupDB()
	defer sqldb.Close()

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

	e.GET("/waiting-opponent/:gameid", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/waiting", map[string]interface{}{
			"JoinGameURL": &url.URL{
				Scheme: c.Scheme(),
				Host:   c.Request().Host,
				Path:   c.Echo().Reverse("joinGame", c.Param("gameid")),
			},
		})
	}).Name = "waitingOpponent"

	e.GET("/join-game/:gameid", func(c echo.Context) error {
		// NOTE: check that no oppoenent is in the game
		gameid := c.Param("gameid")

		if !db.GetOpponentReady(sqldb, ctx, gameid) {
			err := db.UpdateOpponentReady(sqldb, ctx, gameid, true)
			if err != nil {
				c.Logger().Error(err)
			}
			return c.Redirect(http.StatusPermanentRedirect, c.Echo().Reverse("game", c.Param("gameid")))
		} else {
			return c.Redirect(http.StatusPermanentRedirect, c.Echo().Reverse("gameFull"))
		}
	}).Name = "joinGame"

	e.GET("/game-full", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/gamefull", map[string]interface{}{})
	}).Name = "gameFull"

	e.GET("/game/:gameid", func(c echo.Context) error {
		return c.Render(http.StatusOK, "views/game", map[string]interface{}{})
	}).Name = "game"
	// TODO: write the JS websocket client using HTMX
	e.GET("/ws/:id", wsGame).Name = "gameWebsocket"

	e.Logger.Fatal(e.Start(":8000"))
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func createNewGame(c echo.Context, sqldb *sql.DB, ctx context.Context) error {
	gameid := uuid.NewString()

	err := db.CreateEmptyGame(sqldb, ctx, gameid)
	if err != nil {
		c.Logger().Fatal(err)
	}
	res := db.GetGame(sqldb, ctx, gameid)
	c.Logger().Printf("%v", res)

	waitingOpponentURL := c.Echo().Reverse("waitingOpponent", gameid)
	return c.Redirect(http.StatusPermanentRedirect, waitingOpponentURL)
}

func joinGame(c echo.Context) error {
	// gameid := c.Param("gameid")
	// TODO: get the game id from the url
	// TODO: see if it correspond to an existing game in the redis db
	return nil
}

func wsGame(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("Websocket upgrade error: ", err)
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
		_, _, err = ws.ReadMessage()
		// _, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)

		}
		// c.Logger().Printf("%s\n", msg)
	}
}
