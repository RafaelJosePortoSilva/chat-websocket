package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Verificação que libera todos os usuarios
	},
}

// Estrutura da mensagem
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// Forma de armazenar os dados
var clients = make(map[*websocket.Conn]bool)

// Channel para transporte dos dados pela gorotine
var broadcast = make(chan struct {
	Message Message
	Sender  *websocket.Conn
})

// Main
func main() {

	// Setup das rotas
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", handleConnections)

	// Go Routine com loop do websocket
	go handleMessages()

	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error starting server - " + err.Error())
	}

}

// Home Page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to the Chat Room!")
}

// Loop do websocket
func handleConnections(w http.ResponseWriter, r *http.Request) {

	// Upgrade do status da conexão
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	clients[conn] = true

	// loop
	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			delete(clients, conn)
			return
		}

		// Passa os dados para o channel
		broadcast <- struct {
			Message Message
			Sender  *websocket.Conn
		}{msg, conn}
	}

}

// Recebe os dados do broadcast
func handleMessages() {
	for {
		data := <-broadcast
		msg := data.Message
		sender := data.Sender

		for client := range clients {
			if client != sender {
				err := client.WriteJSON(msg)
				if err != nil {
					fmt.Println(err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
