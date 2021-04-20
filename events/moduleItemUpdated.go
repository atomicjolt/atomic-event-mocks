package events

type ModuleItemUpdated struct {
	ModuleItemCreated
}

func MockModuleItemUpdated() ModuleItemUpdated {
	var mock ModuleItemUpdated
	mock.ModuleItemCreated = MockModuleItemCreated()
	return mock
}
