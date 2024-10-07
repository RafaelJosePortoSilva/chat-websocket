package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"chat-websocket/services"
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