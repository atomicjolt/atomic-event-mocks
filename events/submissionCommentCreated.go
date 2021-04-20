package events

import (
	"github.com/brianvoe/gofakeit/v6"
	"strconv"
	"time"
)

type SubmissionCommentCreated struct {
	Metadata `json:"metadata" fake:"skip"`

	Body struct {
		AttachmentIds       []string  `json:"attachment_ids" fake:"skip"`
		Body                string    `json:"body" fake:"{loremipsumsentence:30}"`
		CreatedAt           time.Time `json:"created_at" fake:"{date}"`
		SubmissionCommentId string    `json:"submission_comment_id" fake:"{number:1,10}"`
		SubmissionId        string    `json:"submission_id" fake:"{number:1,10}"`
		UserId              string    `json:"user_id" fake:"{number:1,10}"`
	} `json:"body"`
}

func MockSubmissionCommentCreated() SubmissionCommentCreated {
	var mock SubmissionCommentCreated
	mock.Metadata = mockMetaData("submission_comment_created")
	gofakeit.Struct(&mock.Body)

	mock.Body.AttachmentIds = make([]string, gofakeit.Number(0, 3))
	for i := 0; i < len(mock.Body.AttachmentIds); i++ {
		mock.Body.AttachmentIds[i] = strconv.Itoa(gofakeit.Number(1, 100))
	}

	return mock
}
