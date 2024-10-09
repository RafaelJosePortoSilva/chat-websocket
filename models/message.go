package models

// Estrutura da mensagem
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	IsSender bool   `json:"issender"`
}

type Message2 struct {
	IDUser  string `json:"IDUser"`
	IDConv  string `json:"IDConv"`
	Message string `json:"message"`
}
