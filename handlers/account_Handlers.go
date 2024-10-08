package handlers

import (
	"chat-websocket/models"
	"chat-websocket/services"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleAuth(w http.ResponseWriter, r *http.Request) {

	var accReq models.Account
	err := json.NewDecoder(r.Body).Decode(&accReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Requisição inválida"}`))
		return
	}

	// Verifica os parâmetros username e password
	user := accReq.Username
	pass := accReq.Password

	auth, id := services.AuthAccount(&user, &pass)
	if !auth {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]bool{"message": false})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Successful Login",
		"userID":  id,
	})

}

func HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler create account")

	// Decodifica o JSON do corpo da requisição
	var accReq models.Account
	err := json.NewDecoder(r.Body).Decode(&accReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Requisição inválida"}`))
		return
	}

	// Verifica os parâmetros username e password
	user := accReq.Username
	pass := accReq.Password

	if user == "" || pass == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Missing data"}`))
		return
	}

	_, err = services.NewAccount(user, pass)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"message": "Error to create account"}`))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Account has been created"}`))
	}
}
