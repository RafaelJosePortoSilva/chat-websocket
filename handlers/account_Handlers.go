package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"chat-websocket/services"
	"fmt"
	"chat-websocket/models"
	"encoding/json"
)

func HandleAuth(w http.ResponseWriter, r *http.Request){

	// Puxa os parâmetros da request
	params := mux.Vars(r)

	user := params["username"]
	pass := params["password"]

	auth := services.AuthAccount(&user, &pass)
	if (auth == true){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": true}`))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": false}`))
	}



}

func HandleCreateAccount(w http.ResponseWriter, r *http.Request){
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
		w.Write([]byte(`{"message": "Faltam dados"}`))
		return
	}

	_, err = services.NewAccount(user, pass)
	if( err != nil) {
		fmt.Println(err)
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"message": "Error to create account"}`))
	} else{
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Account has been created"}`))
	}
}
