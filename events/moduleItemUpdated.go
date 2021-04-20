package events

type ModuleItemUpdated struct {
	ModuleItemCreated
}

func (mock *ModuleItemUpdated) Mock() {
	mock.ModuleItemCreated = ModuleItemCreated{}
	mock.ModuleItemCreated.Mock()
	mock.Metadata.EventName = "module_item_updated"
}
