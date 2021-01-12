package server

import (
	"github.com/gorilla/websocket"
	"net/http"

	"github.com/aaorlov/stream/log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// HandleWS serve ws
func HandleWS(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("WS connection upgrade failing: %s\n", err)
		return
	}

	log.Infof("%s", ws)

}
