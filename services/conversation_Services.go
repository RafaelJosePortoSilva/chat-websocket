package services

import (
	"chat-websocket/models"
	"strconv"
)

var Conversations = make(map[string]*models.Conversation)

func SendMessagesToUsers(msg *models.Message, convID string) {

	conversation := FetchConversationById(convID)

	for _, user := range conversation.Clients {
		BroadcastMessage(*msg, user.WsConnection)
	}

}

func CreateConversation(title string) {
	newConv := models.Conversation{
		Messages: nil,
		Clients:  nil,
		Title:    title,
		ID:       generateConversationID(),
	}
	Conversations[newConv.ID] = &newConv
}

func FetchConversationById(convID string) *models.Conversation {
	return Conversations[convID]
}

func AddUserToConversation(userID string, convID string) {
	conv := FetchConversationById(convID)
	user := FetchUser(userID)
	if conv != nil && user != nil {
		conv.Clients = append(conv.Clients, *user)
	}
}

func generateConversationID() string {
	return strconv.Itoa(len(Conversations) + 1)
}
