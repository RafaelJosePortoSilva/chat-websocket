package services

import (
	"chat-websocket/models"
	"fmt"
)

var Users = make(map[string]*models.User)

func NewUser(IdUser string) {

	Users[IdUser] = &models.User{
		ID:          IdUser,
		Name:        "",
		PhoneNumber: "",
	}
	return
}

func FetchUser(id string) *models.User {
	return Users[id]
}

func setUserName(id string, name string) {
	user := FetchUser(id)
	if user != nil {
		user.Name = name
		fmt.Printf("Successful")
	} else {
		fmt.Printf("User id=%s not found\n", id)
	}
}

func setUserPhoneNumber(id string, phone string) {
	user := FetchUser(id)
	if user != nil {
		user.PhoneNumber = phone
		fmt.Printf("Successful")
	} else {
		fmt.Printf("User id=%s not found\n", id)
	}
}
