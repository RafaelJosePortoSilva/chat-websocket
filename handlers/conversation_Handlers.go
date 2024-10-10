package handlers

import (
	"chat-websocket/models"
	"chat-websocket/services"
	"encoding/json"
	"net/http"
)

func HandleCreateConversation(w http.ResponseWriter, r *http.Request) {

	var conv models.Conversation
	err := json.NewDecoder(r.Body).Decode(&conv)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Requisição inválida"}`))
		return
	}

	if conv.Title != "" {
		services.CreateConversation(conv.Title)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Missing data"}`))
		return
	}

}

func HandleAddUserToConversation(w http.ResponseWriter, r *http.Request) {

	var req models.ReqAddUserToConv

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Requisição inválida"}`))
		return
	}

	services.AddUserToConversation(req.IDUser, req.IDConv)

}

func HandleSendMessagesToUsers(w http.ResponseWriter, r *http.Request) {

	var msg models.Message

	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Requisição inválida"}`))
		return
	}

	services.SendMessagesToUsers(&msg, msg.IDConv)

}
