package models

import "github.com/gorilla/websocket"

// Usado para mandar e receber as mensagens
// Para não expôr as informações de login do Account
type User struct {
	ID           string `json:"ID"`
	PhoneNumber  string `json:"phone_number"`
	Name         string `json:"name"`
	WsConnection *websocket.Conn
}
