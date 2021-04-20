package events

import (
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type AssignmentCreated struct {
	Metadata `json:"metadata" fake:"skip"`

	Body struct {
		AssignmentGroupId               string    `json:"assignment_group_id" fake:"{number:1,100}"`
		AssignmentId                    string    `json:"assignment_id" fake:"{number:1,100}"`
		ContextId                       string    `json:"context_id" fake:"{number:1,100}"`
		ContextType                     string    `json:"context_type" fake:"skip"`
		ContextUuid                     string    `json:"context_uuid" fake:"{uuid}"`
		Description                     string    `json:"description" fake:"{loremipsumsentence:30}"`
		DueAt                           time.Time `json:"due_at" fake:"{date}"`
		LockAt                          time.Time `json:"lock_at" fake:"{date}"`
		LtiAssignmentId                 string    `json:"lti_assignment_id" fake:"{number:1,100}"`
		LtiResourceLinkId               string    `json:"lti_resource_link_id" fake:"{number:1,100}"`
		LtiResourceLinkIdDuplicatedFrom string    `json:"lti_resource_link_id_duplicated_from" fake:"{number:1,100}"`
		PointsPossible                  int       `json:"points_possible" fake:"{number:1,100}"`
		SubmissionTypes                 string    `json:"submission_types" fake:"skip"`
		Title                           string    `json:"title" fake:"{buzzword}"`
		UnlockAt                        time.Time `json:"unlock_at" fake:"{date}"`
		UpdatedAt                       time.Time `json:"updated_at" fake:"{date}"`
		WorkflowState                   string    `json:"workflow_state" fake:"skip"`
	} `json:"body"`
}

func (mock *AssignmentCreated) Mock() {
	mock.Metadata = mockMetaData("assignment_created")
	gofakeit.Struct(&mock.Body)

	mock.Body.ContextType = "Course"
	mock.Body.WorkflowState = "published"

	submissionTypes := []string{
		"discussion_topic",
		"external_tool",
		"media_recording",
		"none",
		"not_graded",
		"online_quiz",
		"online_text_entry",
		"online_upload",
		"online_url",
		"on_paper",
	}
	mock.Body.SubmissionTypes = gofakeit.RandomString(submissionTypes)
}
