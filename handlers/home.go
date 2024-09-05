package handlers

import (
	"fmt"
	"net/http"
)

// Home Page
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Welcome to the Chat Room!")
}
