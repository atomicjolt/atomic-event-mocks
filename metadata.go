package main

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

type Metadata struct {
	RootAccountUuid    string    `json:"root_account_uuid"`
	RootAccountId      string    `json:"root_account_id"`
	RootAccountLtiGuid string    `json:"root_account_lti_guid"`
	UserLogin          string    `json:"user_login"`
	UserAccountId      string    `json:"user_account_id"`
	UserSisId          string    `json:"user_sis_id"`
	UserId             string    `json:"user_id"`
	TimeZone           string    `json:"time_zone"`
	ContextType        string    `json:"context_type"`
	ContextId          string    `json:"context_id"`
	ContextSisSourceId string    `json:"context_sis_source_id"`
	ContextAccountId   string    `json:"context_account_id"`
	ContextRole        string    `json:"context_role"`
	RequestId          string    `json:"request_id"`
	SessionId          string    `json:"session_id"`
	Hostname           string    `json:"hostname"`
	HttpMethod         string    `json:"http_method"`
	UserAgent          string    `json:"user_agent"`
	ClientIp           string    `json:"client_ip"`
	Url                string    `json:"url"`
	Referrer           string    `json:"referrer"`
	Producer           string    `json:"producer"`
	EventName          string    `json:"event_name"`
	EventTime          time.Time `json:"event_time"`
}

func mockMetaData() Metadata {
	var m Metadata

	json.Unmarshal([]byte(`{
    "root_account_uuid": "44fJ44GgJ29gJBsl43JLKgljsBIOTsbnKT48932g",
    "root_account_id": "10000000000001",
    "root_account_lti_guid": "794d72b707af6ea82cfe3d5d473f16888a8366c7.canvas.docker",
    "user_login": "oxana@instructure.com",
    "user_account_id": "10000000000002",
    "user_sis_id": null,
    "user_id": "21070000000000001",
    "time_zone": "America/Denver",
    "context_type": "Course",
    "context_id": "21070000000000002",
    "context_sis_source_id": "194387",
    "context_account_id": "21070000000000003",
    "context_role": "TeacherEnrollment",
    "request_id": "98e1b771-fe22-4481-8264-d523dadb16b1",
    "session_id": "242872453a9d69f7ccddeb4788d22506",
    "hostname": "oxana.instructure.com",
    "http_method": "POST",
    "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36",
    "client_ip": "93.184.216.34",
    "url": "http://oxana.instructure.com/courses/2/gradebook/update_submission",
    "referrer": "http://oxana.instructure.com/courses/2/gradebook/speed_grader?assignment_id=39&student_id=2",
    "producer": "canvas",
    "event_name": "course_grade_change",
    "event_time": "2019-12-11T16:26:34.552Z"
	}`), &m)

	m.EventTime = gofakeit.DateRange(time.Now().AddDate(0, 0, -1), time.Now())
	return m
}
