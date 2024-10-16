package main

import (
	"chat-websocket/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Main
func main() {

	r := mux.NewRouter()

	// Serve arquivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Setup das rotas
	r.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/chat.html")
	})
	r.HandleFunc("/chat/ws", handlers.HandleSendMessagesToUsers).Methods("GET")
	//r.HandleFunc("/chat/{IDConv}/ws", handlers.HandleConnections).Methods("GET")

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/login.html")
	})
	r.HandleFunc("/login/auth", handlers.HandleAuth).Methods("POST")
	r.HandleFunc("/login/newauth", handlers.HandleCreateAccount).Methods("POST")

	r.HandleFunc("/conversation", handlers.HandleCreateConversation).Methods("POST")
	r.HandleFunc("/conversation/user", handlers.HandleAddUserToConversation).Methods("POST")

	// Go Routine com loop do websocket
	go handlers.HandleMessages()

	fmt.Println("Server started on :8080")
	err := http.ListenAndServe("0.0.0.0:8080", r)
	if err != nil {
		panic("Error starting server - " + err.Error())
	}

}
