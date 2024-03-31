package handler

import (
	"fmt"
	"log"
	"net/http"
	"sketchNow_service/lib"

	"github.com/google/uuid"
)

func WebsocketHander(hub *lib.Hub, w http.ResponseWriter, r *http.Request, boardRoomId uuid.UUID) {
	conn, err := lib.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &lib.Client{Hub: hub, Conn: conn, BoardRoomId: boardRoomId, Send: make(chan []byte, 256)}
	client.Hub.Register <- client
	
	fmt.Println("Websocket connected", client.Conn.RemoteAddr())
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}