package main

import (
	"chat-websocket/handlers"
	"fmt"
	"net/http"
)

// Main
func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Setup das rotas
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/chat.html")
	})
	http.HandleFunc("/auth", handlers.HandleAuth)
	http.HandleFunc("/ws", handlers.HandleConnections)

	// Go Routine com loop do websocket
	go handlers.HandleMessages()

	fmt.Println("Server started on :8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic("Error starting server - " + err.Error())
	}

}
