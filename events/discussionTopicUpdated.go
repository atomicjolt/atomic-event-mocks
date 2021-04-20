package events

type DiscussionTopicUpdated struct {
	DiscussionTopicCreated
}

func (mock *DiscussionTopicUpdated) Mock() {
	mock.DiscussionTopicCreated = DiscussionTopicCreated{}
	mock.DiscussionTopicCreated.Mock()
	mock.Metadata.EventName = "discussion_topic_updated"
}
