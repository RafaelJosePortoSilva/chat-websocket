package models

type Conversation struct {
	Messages []Message `json:"messages"`
	Clients  []Account `json:"clients"`
}
