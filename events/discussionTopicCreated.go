package events

import (
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type DiscussionTopicCreated struct {
	Metadata `json:"metadata" fake:"skip"`

	Body struct {
		AssignmentId      string    `json:"assignment_id" fake:"{number:1,100}"`
		Body              string    `json:"body" fake:"{loremipsumsentence:30}"`
		ContextId         string    `json:"context_id" fake:"{number:1,100}"`
		ContextType       string    `json:"context_type" fake:"skip"`
		DiscussionTopicId string    `json:"discussion_topic_id" fake:"{number:1,100}"`
		IsAnnouncement    bool      `json:"is_announcement" fake:"{bool}"`
		LockAt            time.Time `json:"lock_at" fake:"{date}"`
		Title             string    `json:"title" fake:"{buzzword}"`
		UpdatedAt         time.Time `json:"updated_at" fake:"{date}"`
		WorkflowState     string    `json:"workflow_state" fake:"skip"`
	} `json:"body"`
}

func MockDiscussionTopicCreated() DiscussionTopicCreated {
	var mock DiscussionTopicCreated
	mock.Metadata = mockMetaData("discussion_entry_created")
	gofakeit.Struct(&mock.Body)

	mock.Body.ContextType = "Course"
	mock.Body.WorkflowState = "active"

	return mock
}
