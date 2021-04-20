package events

type GroupMembershipCreated struct {
	Metadata `json:"metadata"`

	Body struct {
		GroupCategoryId   string `json:"group_category_id" fake:"{number:1,100}"`
		GroupCategoryName string `json:"group_category_name" fake:"{buzzword}"`
		GroupId           string `json:"group_id" fake:"{number:1,100}"`
		GroupMembershipId string `json:"group_membership_id" fake:"{number:1,100}"`
		GroupName         string `json:"group_name" fake:"{buzzword}"`
		UserId            string `json:"user_id" fake:"{number:1,100}"`
		WorkflowState     string `json:"workflow_state" fake:"skip"`
	} `json:"body"`
}

func (mock *GroupMembershipCreated) Mock() {
	mock.Metadata = mockMetaData("group_membership_created")
	mock.Body.WorkflowState = "active"
}
