package events

import (
	"github.com/brianvoe/gofakeit/v6"
)

type AssignmentUpdated struct {
	AssignmentCreated
}

func MockAssignmentUpdated() AssignmentUpdated {
	var mock AssignmentUpdated
	mock.Metadata = mockMetaData("assignment_updated")
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

	return mock
}
