package events

type LoggedOut struct {
	Metadata `json:"metadata"`
	Body     struct{} `json:"body"`
}

func MockLoggedOut() LoggedOut {
	var mock LoggedOut
	mock.Metadata = mockMetaData("logged_out")

	return mock
}
