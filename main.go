package main

import (
	"chat-websocket/handlers"
	"fmt"
	"net/http"
)

// Main
func main() {

	// Setup das rotas
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", handlers.HandleConnections)

	// Go Routine com loop do websocket
	go handlers.HandleMessages()

	fmt.Println("Server started on :8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic("Error starting server - " + err.Error())
	}

}
