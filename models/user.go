package models

// Usado para mandar e receber as mensagens
// Para não expôr as informações de login do Account
type User struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}
