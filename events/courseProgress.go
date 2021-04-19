package events

import (
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type CourseProgress struct {
	Metadata `json:"metadata" fake:"skip"`
	Body     struct {
		Course struct {
			AccountId   string `json:"account_id" fake:"{number:1,100}"`
			Id          string `json:"id" fake:"{number:1,100}"`
			Name        string `json:"name" fake:"{buzzword}"`
			SisSourceId string `json:"sis_source_id" fake:"{number:1,100}"`
		} `json:"course"`
		Progress struct {
			CompletedAt               time.Time `json:"completed_at" fake:"{date}"`
			NextRequirementUrl        string    `json:"next_requirement_url" fake:"skip"`
			RequirementCompletedCount int       `json:"requirement_completed_count" fake:"{number:1,100}"`
			RequirementCount          int       `json:"requirement_count" fake:"{number:1,100}"`
		} `json:"progress"`
		User struct {
			Email string `json:"email" fake:"{email}"`
			Id    string `json:"id" fake:"{number:1,10}"`
			Name  string `json:"name" fake:"{name}"`
		} `json:"user"`
	} `json:"body"`
}

func MockCourseProgress() CourseProgress {
	var mock CourseProgress
	mock.Metadata = mockMetaData("course_grade_change")
	gofakeit.Struct(&mock.Body)

	mock.Body.Progress.NextRequirementUrl = "http:/oxana.instructure.com/courses/1234567/modules/items/12345"

	return mock
}
