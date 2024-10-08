package services

import (
	"chat-websocket/models"
)

var Users = make(map[string]*models.User)

func NewUser(IdUser string) {

	newUs := models.User{
		ID:          IdUser,
		Name:        "",
		PhoneNumber: "",
	}
	Users[IdUser] = &newUs
	return
}

func FetchUser(id string) *models.User {
	return Users[id]
}
