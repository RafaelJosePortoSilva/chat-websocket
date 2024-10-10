package models

// Estrutura da mensagem
type Message struct {
	IDUser  string `json:"IDUser"`
	IDConv  string `json:"IDConv"`
	Message string `json:"message"`
}
