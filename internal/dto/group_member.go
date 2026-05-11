package dto

type GroupMemberRequest struct {
	GroupId string `json:"group_id"`
	UserId  string `json:"user_id"`
	Email   string `json:"email"`
}

type GroupMemberResponse struct {
	Id      string `json:"id,omitempty"`
	GroupId string `json:"group_id,omitempty"`
	UserId  string `json:"user_id,omitempty"`
	Error   string `json:"error,omitempty"`
}

type GroupMemberListResponse struct {
	GroupMembers []GroupMemberResponse `json:"group_members,omitempty"`
	Error        string                `json:"error,omitempty"`
}
