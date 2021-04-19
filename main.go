package main

import (
	"atomicjolt.com/MockCanvasEvents/events"
	"encoding/json"
	"os"
)

func main() {
	m := events.MockAssignmentUpdated()

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")

	if err := enc.Encode(&m); err != nil {
		println(err)
	}
}
