package events

type AssignmentUpdated struct {
	AssignmentCreated
}

func MockAssignmentUpdated() AssignmentUpdated {
	var mock AssignmentUpdated
	mock.AssignmentCreated = MockAssignmentCreated()
	return mock
}
