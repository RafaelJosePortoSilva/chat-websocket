package models

// Estrutura da mensagem
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	IsSender bool   `json:"issender"`
}
