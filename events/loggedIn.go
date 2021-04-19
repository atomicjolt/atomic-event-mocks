package events

import (
	"github.com/brianvoe/gofakeit/v6"
)

type LoggedIn struct {
	Metadata `json:"metadata" fake:"skip"`
	Body     struct {
		RedirectUrl string `json:"redirect_url" fake:"skip"`
	} `json:"body"`
}

func MockLoggedIn() LoggedIn {
	var mock LoggedIn
	mock.Metadata = mockMetaData("logged_in")
	gofakeit.Struct(&mock.Body)

	mock.Body.RedirectUrl = "https://oxana.instructure.com/"

	return mock
}
