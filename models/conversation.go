package models

type Conversation struct {
	Messages []Message `json:"messages"`
	Clients  []User    `json:"clients"`
	Title    string    `json:"title"`
	ID       string    `json:"IDConv"`
}

type ReqAddUserToConv struct {
	IDUser string `json:"IDUser"`
	IDConv string `json:"IDConv"`
}
