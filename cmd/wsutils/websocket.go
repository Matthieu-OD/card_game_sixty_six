package wsutils

import (
	"github.com/gorilla/websocket"
)

// websocket variables
var (
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)
