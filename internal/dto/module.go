package dto

type ModuleRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Position    int64  `json:"position"`
	CourseId    string `json:"course_id"`
}

type ModuleResponse struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Position    int64  `json:"position,omitempty"`
	CourseId    string `json:"course_id,omitempty"`
	Error       string `json:"error,omitempty"`
}

type ModuleListResponse struct {
	Modules []ModuleResponse `json:"modules,omitempty"`
	Error   string           `json:"error,omitempty"`
}
