package models

type Conversation struct {
	Messages []Message `json:"messages"`
	Clients  []User    `json:"clients"`
	Title    string
	ID       string
}
