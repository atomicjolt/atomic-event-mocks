package events

import (
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type LearningOutcomeResultCreated struct {
	Metadata `json:"metadata" fake:"skip"`
	Body     struct {
		AssessedAt        time.Time `json:"assessed_at" fake:"{date}"`
		Attempt           int       `json:"attempt" fake:"{number:1,10}"`
		CreatedAt         time.Time `json:"created_at"  fake:"{date}"`
		LearningOutcomeId string    `json:"learning_outcome_id" fake:"{number:1,10}"`
		Mastery           bool      `json:"mastery" fake:"{bool}"`
		OriginalMastery   bool      `json:"original_mastery" fake:"{bool}"`
		OriginalPossible  int       `json:"original_possible" fake:"{number:1,10}"`
		OriginalScore     int       `json:"original_score" fake:"{number:1,10}"`
		Percent           int       `json:"percent" fake:"{number:1,100}"`
		Possible          int       `json:"possible" fake:"{number:1,10}"`
		Score             int       `json:"score" fake:"{number:1,10}"`
		Title             string    `json:"title" fake:"{buzzword}"`
	} `json:"body"`
}

func MockLearningOutcomeResultCreated() LearningOutcomeResultCreated {
	var mock LearningOutcomeResultCreated
	mock.Metadata = mockMetaData("learning_outcome_result_created")
	gofakeit.Struct(&mock.Body)
	return mock
}
