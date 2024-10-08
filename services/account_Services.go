package services

import (
	"fmt"
)

// Procurar hash username e retornar senha
// Não é feito para ser seguro
var Accounts = make(map[string]string)

func NewAccount(username string, password string) (*string, error) {
	fmt.Println(username)
	if FetchAccount(&username) != "" {
		return &username, fmt.Errorf("User %s already exists", username)
	} else {
		Accounts[username] = password
		return &username, nil
	}
}

func FetchAccount(username *string) string {
	fmt.Printf("Conta: %s\n", Accounts[*username])
	if Accounts[*username] != "" {
		return Accounts[*username]
	} else {
		return ""
	}
}

func AuthAccount(username *string, password *string) bool {
	if FetchAccount(username) != "" && *password != "" {
		return (Accounts[*username] == *password)
	} else {
		fmt.Printf("User %s doesn't exists or the password is invalid.", *username)
		return false
	}
}
