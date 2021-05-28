package lib

import (
	"fmt"
	"github.com/atomicjolt/atomic-event-mocks/events"
)

func InstanceFromName(name string) (events.Event, error) {
	switch name {
	case "submission_created":
		return new(events.SubmissionCreated), nil
	case "syllabus_updated":
		return new(events.SyllabusUpdated), nil
	case "submission_comment_created":
		return new(events.SubmissionCommentCreated), nil
	case "module_item_created":
		return new(events.ModuleItemCreated), nil
	case "logged_out":
		return new(events.LoggedOut), nil
	case "logged_in":
		return new(events.LoggedIn), nil
	case "learning_outcome_result_updated":
		return new(events.LearningOutcomeResultUpdated), nil
	case "learning_outcome_result_created":
		return new(events.LearningOutcomeResultCreated), nil
	case "group_membership_created":
		return new(events.GroupMembershipCreated), nil
	case "group_created":
		return new(events.GroupCreated), nil
	case "grade_change":
		return new(events.GradeChange), nil
	case "discussion_topic_created":
		return new(events.DiscussionTopicCreated), nil
	case "discussion_entry_submitted":
		return new(events.DiscussionEntrySubmitted), nil
	case "discussion_entry_created":
		return new(events.DiscussionEntryCreated), nil
	case "course_progress":
		return new(events.CourseProgress), nil
	case "course_grade_change":
		return new(events.CourseGradeChange), nil
	case "conversation_message_created":
		return new(events.ConversationMessageCreated), nil
	case "asset_accessed":
		return new(events.AssetAccessed), nil
	case "assignment_created":
		return new(events.AssignmentCreated), nil
	case "discussion_topic_updated":
		return new(events.DiscussionTopicUpdated), nil
	case "assignment_updated":
		return new(events.AssignmentUpdated), nil
	case "module_item_updated":
		return new(events.ModuleItemUpdated), nil
	default:
		return nil, fmt.Errorf("unknown event name: %s", name)
	}
}
