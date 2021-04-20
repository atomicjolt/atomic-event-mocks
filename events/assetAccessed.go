package events

import (
	"github.com/brianvoe/gofakeit/v6"
)

type AssetAccessed struct {
	Metadata `json:"metadata" fake:"skip"`
	Body     struct {
		AssetName    string `json:"asset_name" fake:"{buzzword}"`
		AssetType    string `json:"asset_type" fake:"skip"`
		AssetId      string `json:"asset_id" fake:"{number:1,10}"`
		AssetSubtype string `json:"asset_subtype" fake:"{buzzword}"`
		Category     string `json:"category" fake:"{buzzword}"`
		Role         string `json:"role" fake:"{buzzword}"`
		Level        string `json:"level" fake:"{buzzword}"`
	} `json:"body"`
}

func (mock *AssetAccessed) Mock() {
	assetTypes := []string{
		"group",
		"course",
		"enrollment",
		"user",
		"collaboration",
		"account",
		"quizzes:quiz",
		"lti/tool_proxy",
		"assignment",
		"calendar_event",
		"web_conference",
		"learning_outcome",
		"wiki_page",
		"discussion_topic",
		"attachment",
	}

	mock.Metadata = mockMetaData("asset_accessed")
	gofakeit.Struct(&mock.Body)

	mock.Body.AssetType = gofakeit.RandomString(assetTypes)
}
