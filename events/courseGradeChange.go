package events

import (
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type CourseGradeChange struct {
	Metadata `json:"metadata" fake:"skip"`
	Body     struct {
		UserId                  string    `json:"user_id" fake:"{number:1,10}"`
		CourseId                string    `json:"course_id" fake:"{number:1,10}"`
		WorkflowState           string    `json:"workflow_state" fake:"{randomstring:[active]}"` //Could add "deleted"
		CreatedAt               time.Time `json:"created_at" fake:"{date}"`
		UpdatedAt               time.Time `json:"updated_at" fake:"{date}"`
		CurrentScore            float64   `json:"current_score" fake:"{number:1,100}"`
		OldCurrentScore         float64   `json:"old_current_score" fake:"{number:1,100}"`
		FinalScore              float64   `json:"final_score" fake:"{number:1,100}"`
		OldFinalScore           float64   `json:"old_final_score" fake:"{number:1,100}"`
		UnpostedCurrentScore    float64   `json:"unposted_current_score" fake:"{number:1,100}"`
		OldUnpostedCurrentScore float64   `json:"old_unposted_current_score" fake:"{number:1,100}"`
		UnpostedFinalScore      float64   `json:"unposted_final_score" fake:"{number:1,100}"`
		OldUnpostedFinalScore   float64   `json:"old_unposted_final_score" fake:"{number:1,100}"`
	} `json:"body"`
}

func (mock *CourseGradeChange) Mock() {
	mock.Metadata = mockMetaData("course_grade_change")
	gofakeit.Struct(&mock.Body)
}
