package events

import (
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type ConversationMessageCreated struct {
	Metadata `json:"metadata" fake:"skip"`

	Body struct {
		AuthorId       string    `json:"author_id" fake:"{number:1,100}"`
		ConversationId string    `json:"conversation_id" fake:"{number:1,100}"`
		CreatedAt      time.Time `json:"created_at" fake:"{date}"`
		MessageId      string    `json:"message_id" fake:"{number:1,100}"`
	} `json:"body"`
}

func MockConversationMessageCreated() ConversationMessageCreated {
	var mock ConversationMessageCreated
	mock.Metadata = mockMetaData("conversation_message_created")
	gofakeit.Struct(&mock.Body)

	return mock
}
