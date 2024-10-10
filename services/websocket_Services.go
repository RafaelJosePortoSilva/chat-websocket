package services

import (
	"chat-websocket/models"
	"fmt"

	"github.com/gorilla/websocket"
)

// Forma de armazenar os dados
var Clients = make(map[*websocket.Conn]bool)

// Channel para transporte dos dados pela gorotine
var Broadcast = make(chan struct {
	Message  models.Message
	Receiver *websocket.Conn
})

func RegisterClient(conn *websocket.Conn) {
	Clients[conn] = true
	fmt.Println("Nova Conexão")
}

func DeleteClient(conn *websocket.Conn) {
	delete(Clients, conn)
}

func BroadcastMessage(msg models.Message, receiver *websocket.Conn) {
	// Passa os dados para o channel
	Broadcast <- struct {
		Message  models.Message
		Receiver *websocket.Conn
	}{msg, receiver}
}
