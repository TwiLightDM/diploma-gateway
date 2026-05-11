package dto

type TaskAttemptAnswerRequest struct {
	TaskId            string   `json:"task_id"`
	TextAnswer        string   `json:"text_answer"`
	SelectedOptionIds []string `json:"selected_option_ids"`
}

type TaskAttemptRequest struct {
	UserId string `json:"user_id"`

	CourseId string `json:"course_id"`
	ModuleId string `json:"module_id"`

	Answers []TaskAttemptAnswerRequest `json:"answers"`
}

type TaskAttemptAnswerResponse struct {
	TaskId            string   `json:"task_id"`
	TextAnswer        string   `json:"text_answer"`
	SelectedOptionIds []string `json:"selected_option_ids"`
	IsCorrect         bool     `json:"is_correct"`
}

type TaskAttemptResponse struct {
	Id string `json:"id"`

	UserId string `json:"user_id"`

	CourseId string `json:"course_id"`
	ModuleId string `json:"module_id"`

	AttemptNumber int `json:"attempt_number"`

	Answers []TaskAttemptAnswerResponse `json:"answers"`

	CorrectAnswers int     `json:"correct_answers"`
	TotalQuestions int     `json:"total_questions"`
	Score          float64 `json:"score"`
}

type TaskAttemptListResponse struct {
	TaskAttempts []TaskAttemptResponse `json:"task_attempts"`
}
