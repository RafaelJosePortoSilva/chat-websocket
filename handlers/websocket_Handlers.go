package handlers

import (
	"chat-websocket/models"
	"chat-websocket/services"

	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Verificação que libera todos os usuarios
	},
}

// Loop do websocket
func HandleConnections(w http.ResponseWriter, r *http.Request) {

	// Upgrade do status da conexão
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	services.RegisterClient(conn)

	// loop
	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			services.DeleteClient(conn)
			return
		}
	}

}

// Recebe os dados do broadcast
func HandleMessages() {
	for {
		data := <-services.Broadcast
		msg := data.Message
		sender := data.Sender

		for client := range services.Clients {
			if client != sender {
				err := client.WriteJSON(msg)
				if err != nil {
					fmt.Println(err)
					client.Close()
					services.DeleteClient(client)
				}
			}
		}
	}
}
