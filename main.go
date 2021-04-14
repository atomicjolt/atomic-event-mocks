package main

import (
	"encoding/json"
	"os"
)

func main() {
	m := mockCourseGradeChange()

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(&m); err != nil {
		println(err)
	}
}
