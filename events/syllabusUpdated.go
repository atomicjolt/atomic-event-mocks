package events

import (
	"github.com/brianvoe/gofakeit/v6"
)

type SyllabusUpdated struct {
	Metadata `json:"metadata" fake:"skip"`

	Body struct {
		CourseId        string `json:"course_id" fake:"{number:1,100}"`
		OldSyllabusBody string `json:"old_syllabus_body" fake:"{loremipsumsentence:30}"`
		SyllabusBody    string `json:"syllabus_body" fake:"{loremipsumsentence:30}"`
	} `json:"body"`
}

func (mock *SyllabusUpdated) Mock() {
	mock.Metadata = mockMetaData("syllabus_updated")
	gofakeit.Struct(&mock.Body)
}
