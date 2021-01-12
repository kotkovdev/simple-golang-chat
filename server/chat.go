package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Token   string `json:"token"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

var clients = make(map[*websocket.Conn]struct{}) // connected clients
var broadcast = make(chan Message)               // broadcast channel

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()
	clients[ws] = struct{}{}

	for {
		var msg Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			delete(clients, ws)
			break
		}

		broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-broadcast

		log.Println(msg)
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
