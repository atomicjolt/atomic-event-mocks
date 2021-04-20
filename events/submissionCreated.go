package events

import (
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type SubmissionCreated struct {
	Metadata `json:"metadata" fake:"skip"`
	Body     struct {
		AssignmentId    string    `json:"assignment_id" fake:"{number:1,100}"`
		Attempt         int       `json:"attempt" fake:"{number:1,5}"`
		Body            string    `json:"body" fake:"{buzzword}"`
		Grade           string    `json:"grade" fake:"skip"`
		GradedAt        time.Time `json:"graded_at" fake:"{date}"`
		GroupId         string    `json:"group_id" fake:"{number:1,100}"`
		Late            bool      `json:"late" fake:"{bool}"`
		LtiAssignmentId string    `json:"lti_assignment_id" fake:"{number:1,100}"`
		LtiUserId       string    `json:"lti_user_id" fake:"{number:1,100}"`
		Missing         bool      `json:"missing" fake:"{bool}"`
		Score           float64   `json:"score" fake:"{float64range:1,100}"`
		SubmissionId    string    `json:"submission_id" fake:"{number:1,100}"`
		SubmissionType  string    `json:"submission_type" fake:"skip"`
		SubmittedAt     time.Time `json:"submitted_at" fake:"{date}"`
		UpdatedAt       time.Time `json:"updated_at" fake:"{date}"`
		Url             string    `json:"url" fake:"skip"`
		UserId          string    `json:"user_id" fake:"{number:1,100}"`
		WorkflowState   string    `json:"workflow_state" fake:"skip"`
	} `json:"body"`
}

func (mock *SubmissionCreated) Mock() {
	mock.Metadata = mockMetaData("submission_created")
	gofakeit.Struct(&mock.Body)

	grades := []string{"A", "B", "C", "D", "F"}
	mock.Body.Grade = gofakeit.RandomString(grades)

	submissionTypes := []string{
		"basic_lti_launch",
		"discussion_topic",
		"media_recording",
		"online_quiz",
		"online_text_entry",
		"online_upload",
		"online_url",
	}
	mock.Body.SubmissionType = gofakeit.RandomString(submissionTypes)

	mock.Body.Url = "https://test.submission.net"
	mock.Body.WorkflowState = "submitted"
}
