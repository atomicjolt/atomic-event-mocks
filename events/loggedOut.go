package events

type LoggedOut struct {
	Metadata `json:"metadata"`
	Body     struct{} `json:"body"`
}

func (mock *LoggedOut) Mock() {
	mock.Metadata = mockMetaData("logged_out")
}
