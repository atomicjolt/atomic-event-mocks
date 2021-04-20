package events

type GroupCreated struct {
	Metadata `json:"metadata"`

	Body struct {
		AccountId         string `json:"account_id" fake:"{number:1,100}"`
		ContextId         string `json:"context_id" fake:"{number:1,100}"`
		ContextType       string `json:"context_type" fake:"skip"`
		GroupCategoryId   string `json:"group_category_id" fake:"{number:1,100}"`
		GroupCategoryName string `json:"group_category_name" fake:"{buzzword}"`
		GroupId           string `json:"group_id" fake:"{number:1,100}"`
		GroupName         string `json:"group_name" fake:"{buzzword}"`
		MaxMembership     int    `json:"max_membership" fake:"{number:1,100}"`
		Uuid              string `json:"uuid" fake:"{uuid}"`
		WorkflowState     string `json:"workflow_state" fake:"skip"`
	} `json:"body"`
}

func (mock *GroupCreated) Mock() {
	mock.Metadata = mockMetaData("group_created")

	mock.Body.ContextType = "Course"
	mock.Body.WorkflowState = "active"
}
