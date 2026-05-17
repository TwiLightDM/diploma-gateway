package dto

type CompletedTheoryCourseRequest struct {
	CourseId string `json:"course_id"`
	UserId   string `json:"user_id"`
}

type CompletedTheoryCourseResponse struct {
	UserId   string `json:"user_id"`
	CourseId string `json:"course_id"`
}

type CompletedTheoryCourseListResponse struct {
	CompletedTheoryCourses []CompletedTheoryCourseResponse `json:"completed_theory_courses"`
}
