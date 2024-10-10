package models

// Usado para fazer login
type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IdUser   string `json:"IDUser"`
}
