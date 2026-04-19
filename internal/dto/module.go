package dto

type ModuleRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Position    int64  `json:"position"`
	CourseId    string `json:"course_id"`
}

type ModuleResponse struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	Position        int64  `json:"position"`
	CourseId        string `json:"course_id"`
	AmountOfLessons int    `json:"amount_of_lessons"`
}

type ModuleListResponse struct {
	Modules []ModuleResponse `json:"modules"`
}
