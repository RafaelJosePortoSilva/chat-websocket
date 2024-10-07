package services

import (
	"fmt"
	"chat-websocket/models"
)

// Procurar hash username e retornar senha
// Não é feito para ser seguro
var Accounts = make(map[*string]string)

func NewAccount(username string, password string) *models.Account, error{
	if FetchAccount(&username) != "" {
		Accounts[&username] = password
		return Accounts[&username], nil
	} else {
		return nil, fmt.Errorf("User %s already exists", username)
	}
}

func FetchAccount(username *string) string {
	if Accounts[username] != "" {
		return Accounts[username]
	} else {
		return ""
	}
}

func AuthAccount(username *string, password *string) bool {
	if FetchAccount(username) != "" && *password != "" {
		return (Accounts[username] == *password)
	} else {
		fmt.Printf("User %s doesn't exists or the password is invalid.", *username)
		return false
	}
}
