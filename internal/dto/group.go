package dto

type GroupRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	OwnerId     string `json:"owner_id"`
}

type GroupResponse struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	OwnerId     string `json:"owner_id,omitempty"`
	Error       string `json:"error,omitempty"`
}

type GroupListResponse struct {
	Groups []GroupResponse `json:"groups,omitempty"`
	Error  string          `json:"error,omitempty"`
}
