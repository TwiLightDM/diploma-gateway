package dto

type CourseRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AccessType  string `json:"access_type"`
	OwnerId     string `json:"owner_id"`
}

type CourseResponse struct {
	Id              string `json:"id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	AccessType      string `json:"access_type"`
	PublishedAt     string `json:"published_at"`
	OwnerId         string `json:"owner_id"`
	AmountOfModules int    `json:"amount_of_modules"`
	AmountOfLessons int    `json:"amount_of_lessons"`
}

type CourseListResponse struct {
	Courses []CourseResponse `json:"courses"`
}

type ReadCoursesByGroupIdsRequest struct {
	GroupIds []string `json:"group_ids"`
}
