package dto

type TaskOptionRequest struct {
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type TaskOptionResponse struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	IsCorrect bool   `json:"is_correct"`
}

type TaskRequest struct {
	Title string `json:"title"`

	CourseId string `json:"course_id"`
	ModuleId string `json:"module_id"`

	Type string `json:"type"`

	CorrectTextAnswer string `json:"correct_text_answer"`

	Options []TaskOptionRequest `json:"options"`
}

type TaskResponse struct {
	Id string `json:"id"`

	Title string `json:"title"`

	CourseId string `json:"course_id"`
	ModuleId string `json:"module_id"`

	Type string `json:"type"`

	CorrectTextAnswer string `json:"correct_text_answer"`

	Options []TaskOptionResponse `json:"options"`
}

type TaskListResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}
