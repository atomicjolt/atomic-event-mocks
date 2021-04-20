package events

import (
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type DiscussionEntrySubmitted struct {
	Metadata `json:"metadata" fake:"skip"`

	Body struct {
		AssignmentId            string    `json:"assignment_id" fake:"{number:1,100}"`
		CreatedAt               time.Time `json:"created_at" fake:"{date}"`
		DiscussionEntryId       string    `json:"discussion_entry_id" fake:"{number:1,100}"`
		DiscussionTopicId       string    `json:"discussion_topic_id" fake:"{number:1,100}"`
		ParentDiscussionEntryId string    `json:"parent_discussion_entry_id" fake:"{number:1,100}"`
		SubmissionId            string    `json:"submission_id" fake:"{number:1,100}"`
		Text                    string    `json:"text" fake:"{loremipsumsentence:30}"`
		UserId                  string    `json:"user_id" fake:"{number:1,100}"`
	} `json:"body"`
}

func (mock *DiscussionEntrySubmitted) Mock() {
	mock.Metadata = mockMetaData("discussion_entry_submitted")
	gofakeit.Struct(&mock.Body)
}
