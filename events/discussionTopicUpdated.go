package events

type DiscussionTopicUpdated struct {
	DiscussionTopicCreated
}

func MockDiscussionTopicUpdated() DiscussionTopicUpdated {
	var mock DiscussionTopicUpdated
	mock.DiscussionTopicCreated = MockDiscussionTopicCreated()
	mock.Metadata.EventName = "discussion_topic_updated"
	return mock
}
