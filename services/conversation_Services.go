package services

import (
	"chat-websocket/models"
	"strconv"
)

var Conversations = make(map[string]*models.Conversation)

func SendMessagesToUsers(msg *models.Message, convID string) {

	conversation := FetchConversationById(convID)

	for _, user := range conversation.Clients {
		BroadcastMessage(*msg, &user.WsConnection)
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

func addUserToConversation(convID string, user *models.User) {
	conv := FetchConversationById(convID)
	if conv != nil {
		append(conv.Clients, user)
	}
}

func generateConversationID() string {
	return strconv.Itoa(len(Conversations) + 1)
}
