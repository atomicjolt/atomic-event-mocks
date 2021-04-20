package events

import (
	"github.com/brianvoe/gofakeit/v6"
)

type GradeChange struct {
	Metadata `json:"metadata" fake:"skip"`

	Body struct {
		AssignmentId      string `json:"assignment_id" fake:"{number:1,100}"`
		Grade             string `json:"grade" fake:"{randstring:[A, B, C, D, F]}"`
		GraderId          string `json:"grader_id" fake:"{number:1,100}"`
		GradingComplete   bool   `json:"grading_complete" fake:"{bool}"`
		Muted             bool   `json:"muted" fake:"{bool}"`
		OldGrade          string `json:"old_grade" fake:"{randstring:[A, B, C, D, F]}"`
		OldPointsPossible int    `json:"old_points_possible" fake:"{number:1,100}"`
		OldScore          int    `json:"old_score" fake:"{number:1,100}"`
		PointsPossible    int    `json:"points_possible" fake:"{number:1,100}"`
		Score             int    `json:"score" fake:"{number:1,100}"`
		StudentId         string `json:"student_id" fake:"{number:1,100}"`
		StudentSisId      string `json:"student_sis_id" fake:"{number:1,100}"`
		SubmissionId      string `json:"submission_id" fake:"{number:1,100}"`
		UserId            string `json:"user_id" fake:"{number:1,100}"`
	} `json:"body"`
}

func MockGradeChange() GradeChange {
	var mock GradeChange
	mock.Metadata = mockMetaData("grade_change")
	gofakeit.Struct(&mock.Body)
	return mock
}
