package dto

type CourseRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AccessType  string `json:"access_type"`
	OwnerId     string `json:"owner_id"`
}

type CourseResponse struct {
	Id          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	AccessType  string `json:"access_type,omitempty"`
	PublishedAt string `json:"published_at,omitempty"`
	OwnerId     string `json:"owner_id,omitempty"`
	Error       string `json:"error,omitempty"`
}

type CourseListResponse struct {
	Courses []CourseResponse `json:"courses,omitempty"`
	Error   string           `json:"error,omitempty"`
}

type ReadCoursesByGroupIdsRequest struct {
	GroupIds []string `json:"group_ids"`
}
