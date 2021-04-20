package events

import (
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type DiscussionEntryCreated struct {
	Metadata `json:"metadata" fake:"skip"`
	Body     struct {
		CreatedAt               time.Time `json:"created_at" fake:"{date}"`
		DiscussionEntryId       string    `json:"discussion_entry_id" fake:"{number:1,100}"`
		DiscussionTopicId       string    `json:"discussion_topic_id" fake:"{number:1,100}"`
		ParentDiscussionEntryId string    `json:"parent_discussion_entry_id" fake:"{number:1,100}"`
		Text                    string    `json:"text" fake:"{loremipsumsentence:30}"`
		UserId                  string    `json:"user_id" fake:"{number:1,100}"`
	} `json:"body"`
}

func (mock *DiscussionEntryCreated) Mock() {
	mock.Metadata = mockMetaData("discussion_entry_created")
	gofakeit.Struct(&mock.Body)
}
