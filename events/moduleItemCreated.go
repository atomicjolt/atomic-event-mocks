package events

import (
	"github.com/brianvoe/gofakeit/v6"
)

type ModuleItemCreated struct {
	Metadata `json:"metadata" fake:"skip"`
	Body     struct {
		ContextId     string `json:"context_id" fake:"{number:1,100}"`
		ContextType   string `json:"context_type" fake:"skip"`
		ModuleId      string `json:"module_id" fake:"{number:1,100}"`
		ModuleItemId  string `json:"module_item_id" fake:"{number:1,100}"`
		Position      int    `json:"position" fake:"{number:1,100}"`
		WorkflowState string `json:"workflow_state" fake:"skip"`
	} `json:"body"`
}

func (mock *ModuleItemCreated) Mock() {
	mock.Metadata = mockMetaData("module_item_created")
	gofakeit.Struct(&mock.Body)

	mock.Body.ContextType = "Course"
	mock.Body.WorkflowState = "active"
}
