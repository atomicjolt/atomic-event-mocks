package events

import (
	"github.com/brianvoe/gofakeit/v6"
)

type ModuleItemUpdated struct {
	//Has the same payload as the created event
	ModuleItemCreated
}

func MockModuleItemUpdated() ModuleItemUpdated {
	var mock ModuleItemUpdated
	mock.Metadata = mockMetaData("module_item_updated")
	gofakeit.Struct(&mock.Body)

	mock.Body.ContextType = "Course"
	mock.Body.WorkflowState = "active"

	return mock
}
