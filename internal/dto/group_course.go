package dto

type GroupCourseRequest struct {
	GroupId  string `json:"group_id"`
	CourseId string `json:"course_id"`
}

type GroupCourseResponse struct {
	Id       string `json:"id,omitempty"`
	GroupId  string `json:"group_id,omitempty"`
	CourseId string `json:"course_id,omitempty"`
	Error    string `json:"error,omitempty"`
}

type GroupCourseListResponse struct {
	GroupCourses []GroupCourseResponse `json:"group_courses,omitempty"`
	Error        string                `json:"error,omitempty"`
}
