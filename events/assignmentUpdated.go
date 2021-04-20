package events

type AssignmentUpdated struct {
	AssignmentCreated
}

func MockAssignmentUpdated() AssignmentUpdated {
	var mock AssignmentUpdated
	mock.AssignmentCreated = MockAssignmentCreated()
	mock.Metadata.EventName = "assignment_updated"
	return mock
}
