package events

type ModuleItemUpdated struct {
	ModuleItemCreated
}

func MockModuleItemUpdated() ModuleItemUpdated {
	var mock ModuleItemUpdated
	mock.ModuleItemCreated = MockModuleItemCreated()
	mock.Metadata.EventName = "module_item_updated"
	return mock
}
