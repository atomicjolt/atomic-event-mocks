package events

type AssignmentUpdated struct {
	AssignmentCreated
}

func (mock *AssignmentUpdated) Mock() {
	mock.AssignmentCreated = AssignmentCreated{}
	mock.AssignmentCreated.Mock()
	mock.Metadata.EventName = "assignment_updated"
}
