package handlers

import (
	"chat-websocket/models"
	"chat-websocket/services"
	"encoding/json"

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

	var msg models.Message

	err = json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Requisição inválida"}`))
		return
	}

	services.RegisterClient(conn, msg.IDUser)

	// loop
	for {
		fmt.Println("Passou pelo loop do handleconnections")

		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			services.DeleteClient(conn)
			return
		}
		services.SendMessagesToUsers(&msg, msg.IDConv)
	}

}

// Recebe os dados do broadcast
func HandleMessages() {
	for {
		data := <-services.Broadcast
		msg := data.Message
		client := data.Receiver

		messageToSend := msg

		err := client.WriteJSON(messageToSend)
		if err != nil {
			fmt.Println(err)
			client.Close()
			services.DeleteClient(client)
		}
	}
}
