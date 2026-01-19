package dto

type LessonRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Position    int64  `json:"position"`
	ModuleId    string `json:"module_id"`
}

type LessonResponse struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Content     string `json:"content,omitempty"`
	Position    int64  `json:"position,omitempty"`
	ModuleId    string `json:"module_id,omitempty"`
	Error       string `json:"error,omitempty"`
}

type LessonListResponse struct {
	Lessons []LessonResponse `json:"lessons,omitempty"`
	Error   string           `json:"error,omitempty"`
}
