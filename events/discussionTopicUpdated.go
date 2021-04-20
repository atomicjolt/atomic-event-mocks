package events

type DiscussionTopicUpdated struct {
	DiscussionTopicCreated
}

func MockDiscussionTopicUpdated() DiscussionTopicUpdated {
	var mock DiscussionTopicUpdated
	mock.DiscussionTopicCreated = MockDiscussionTopicCreated()
	return mock
}
