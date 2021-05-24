package main

import (
	"atomicjolt.com/MockCanvasEvents/events"
	"encoding/json"
)

const ApiUrl = "http://localhost:8888/events"
const Secret = "shared_secret"
const FormKey = "events"

func main() {
	e1 := events.DiscussionEntryCreated{}
	e1.Mock()
	e2 := events.DiscussionEntryCreated{}
	e2.Mock()

	mockedEvents := []events.DiscussionEntryCreated{
		e1,
		e2,
	}

	payload, err := json.Marshal(mockedEvents)
	if err != nil {
		panic(err)
	}

	// Post mocked events
	token, err := lib.EventsJwtFrom(payload, Secret)
	if err != nil {
		panic(err)
	}

	_, err = lib.PostFormTo(ApiUrl, FormKey, token)
	if err != nil {
		panic(err)
	}
}
