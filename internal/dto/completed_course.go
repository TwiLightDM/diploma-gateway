package dto

type CompletedCourseRequest struct {
	CourseId string `json:"course_id"`
	UserId   string `json:"user_id"`
}

type CompletedCourseResponse struct {
	UserId   string `json:"user_id"`
	CourseId string `json:"course_id"`
}

type CompletedCourseListResponse struct {
	CompletedCourse []CompletedCourseResponse `json:"completed_courses"`
}
