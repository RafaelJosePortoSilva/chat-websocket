package services

import (
	"chat-websocket/models"
	"fmt"
	"strconv"
)

// Procurar hash username e retornar senha
// Não é feito para ser seguro
var Accounts = make(map[string]*models.Account)

func NewAccount(username string, password string) (*string, error) {
	fmt.Println(username)
	if FetchAccount(&username) != "" {
		return &username, fmt.Errorf("user %s already exists", username)
	} else {

		newAcc := models.Account{
			Username: username,
			Password: password,
			IdUser:   generateUserID(),
		}
		Accounts[username] = &newAcc
		newUs := newUser(newAcc.IdUser)
		return &username, nil
	}
}

func FetchAccount(username *string) string {
	fmt.Printf("Conta: %s\n", Accounts[*username].Username)
	if Accounts[*username] != nil {
		return Accounts[*username].IdUser
	} else {
		return ""
	}
}

func AuthAccount(username *string, password *string) (bool, string) {
	if FetchAccount(username) != "" && *password != "" {
		if Accounts[*username].Password == *password {
			return true, Accounts[*username].IdUser
		} else {
			return false, ""
		}
	} else {
		fmt.Printf("User %s doesn't exists or the password is invalid.", *username)
		return false, ""
	}
}

func generateUserID() string {
	return strconv.Itoa(len(Accounts) + 1)
}
