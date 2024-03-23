package handler

import (
	"log"
	"net/http"
	"sketchNow_service/lib"
)

func WebsocketHander(hub *lib.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := lib.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &lib.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}