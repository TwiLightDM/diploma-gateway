package dto

type LessonProgressRequest struct {
	LessonId string `json:"lesson_id"`
}

type LessonProgressResponse struct {
	UserId   string `json:"user_id"`
	LessonId string `json:"lesson_id"`
}

type LessonProgressListResponse struct {
	Progress []LessonProgressResponse `json:"progress"`
}

type CourseProgressResponse struct {
	CourseId           string   `json:"course_id"`
	TotalLessons       int      `json:"total_lessons"`
	CompletedLessons   int      `json:"completed_lessons"`
	ProgressPercent    float64  `json:"progress_percent"`
	CompletedLessonIds []string `json:"completed_lesson_ids"`
}

type UserCourseProgressResponse struct {
	UserId           string  `json:"user_id"`
	CompletedLessons int     `json:"completed_lessons"`
	TotalLessons     int     `json:"total_lessons"`
	ProgressPercent  float64 `json:"progress_percent"`
	Completed        bool    `json:"completed"`
}

type CourseStatisticsResponse struct {
	CourseId string                       `json:"course_id"`
	Users    []UserCourseProgressResponse `json:"users"`
}

type ModuleProgressResponse struct {
	ModuleId           string   `json:"module_id"`
	TotalLessons       int      `json:"total_lessons"`
	CompletedLessons   int      `json:"completed_lessons"`
	ProgressPercent    float64  `json:"progress_percent"`
	CompletedLessonIds []string `json:"completed_lesson_ids"`
}

type UserModuleProgressResponse struct {
	UserId           string  `json:"user_id"`
	CompletedLessons int     `json:"completed_lessons"`
	TotalLessons     int     `json:"total_lessons"`
	ProgressPercent  float64 `json:"progress_percent"`
	Completed        bool    `json:"completed"`
}

type ModuleStatisticsResponse struct {
	ModuleId string                       `json:"module_id"`
	Users    []UserModuleProgressResponse `json:"users"`
}
