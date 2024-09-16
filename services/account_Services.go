package services

import (
	"fmt"
)

// Procurar hash username e retornar senha
// Não é feito para ser seguro
var Accounts = make(map[*string]string)

func newAccount(username string, password string) {
	if fetchAccount(&username) != "" {
		Accounts[&username] = password
		fmt.Printf("User %s created", username)
	} else {
		fmt.Printf("User %s already exists", username)
	}
}

func fetchAccount(username *string) string {
	if Accounts[username] != "" {
		return Accounts[username]
	} else {
		return ""
	}
}

func authAccount(username *string, password *string) bool {
	if fetchAccount(username) != "" && *password != "" {
		return (Accounts[username] == *password)
	} else {
		fmt.Printf("User %s doesn't exists or the password is invalid.", *username)
		return false
	}
}
