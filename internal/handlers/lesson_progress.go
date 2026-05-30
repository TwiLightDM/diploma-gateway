package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/TwiLightDM/diploma-gateway/internal/dto"
	"github.com/TwiLightDM/diploma-gateway/internal/grpc/course-service"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LessonProgressHandler struct {
	lessonProgressClient *course_service.CourseClient
}

func NewLessonProgressHandler(client *course_service.CourseClient) *LessonProgressHandler {
	return &LessonProgressHandler{lessonProgressClient: client}
}

// CreateLessonProgress
// @Summary Отметить урок как пройденный
// @Tags Lesson Progress
// @Accept json
// @Produce json
// @Param request body dto.LessonProgressRequest true "Lesson ID"
// @Success 200 {object} dto.LessonProgressResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /lesson-progresses [post]
func (h *LessonProgressHandler) CreateLessonProgress(c echo.Context) error {
	var request dto.LessonProgressRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid request",
		})
	}

	userId := c.Get("user_id").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.lessonProgressClient.CreateLessonProgress(ctx, userId, request.LessonId)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.AlreadyExists:
				return c.JSON(http.StatusConflict, dto.ErrorResponse{
					Error: "progress already exists",
				})
			default:
				return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
					Error: err.Error(),
				})
			}
		}

		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.LessonProgressResponse{
		UserId:   response.LessonProgress.UserId,
		LessonId: response.LessonProgress.LessonId,
	})
}

// ReadLessonProgressByUserId
// @Summary Получить прогресс пользователя по урокам
// @Tags Lesson Progress
// @Produce json
// @Success 200 {object} dto.LessonProgressListResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /lesson-progresses [get]
func (h *LessonProgressHandler) ReadLessonProgressByUserId(c echo.Context) error {
	userId := c.Get("user_id").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.lessonProgressClient.ReadLessonProgressByUserId(ctx, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	progress := make([]dto.LessonProgressResponse, 0, len(response.LessonProgresses))
	for _, p := range response.LessonProgresses {
		progress = append(progress, dto.LessonProgressResponse{
			UserId:   p.UserId,
			LessonId: p.LessonId,
		})
	}

	return c.JSON(http.StatusOK, dto.LessonProgressListResponse{
		Progress: progress,
	})
}

// ReadLessonProgressByUserIdAndLessonId
// @Summary Получить прогресс по уроку
// @Tags Lesson Progress
// @Produce json
// @Param lessonId path string true "Lesson ID"
// @Success 200 {object} dto.LessonProgressResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /lesson-progresses/lessons/{lessonId} [get]
func (h *LessonProgressHandler) ReadLessonProgressByUserIdAndLessonId(c echo.Context) error {
	lessonId := c.Param("lessonId")
	userId := c.Get("user_id").(string)

	if lessonId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.lessonProgressClient.ReadLessonProgressByUserIdAndLessonId(ctx, userId, lessonId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.LessonProgressResponse{
		UserId:   response.LessonProgress.UserId,
		LessonId: response.LessonProgress.LessonId,
	})
}

// ReadCourseProgress
// @Summary Получить прогресс курса
// @Tags Course Progress
// @Produce json
// @Param courseId path string true "Course ID"
// @Success 200 {object} dto.CourseProgressResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /progresses/courses/{courseId} [get]
func (h *LessonProgressHandler) ReadCourseProgress(c echo.Context) error {
	courseId := c.Param("courseId")
	userId := c.Get("user_id").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.lessonProgressClient.ReadCourseProgressByUserId(ctx, userId, courseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.CourseProgressResponse{
		CourseId:           response.CourseProgress.CourseId,
		TotalLessons:       int(response.CourseProgress.TotalLessons),
		CompletedLessons:   int(response.CourseProgress.CompletedLessons),
		ProgressPercent:    response.CourseProgress.ProgressPercent,
		CompletedLessonIds: response.CourseProgress.CompletedLessonIds,
	})
}

// ReadCourseStatistics
// @Summary Получить статистику курса
// @Tags Course Progress
// @Produce json
// @Param courseId path string true "Course ID"
// @Success 200 {object} dto.CourseStatisticsResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /progresses/courses/{courseId}/statistics [get]
func (h *LessonProgressHandler) ReadCourseStatistics(c echo.Context) error {
	courseId := c.Param("courseId")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.lessonProgressClient.ReadCourseStatistics(ctx, courseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	users := make([]dto.UserCourseProgressResponse, 0, len(response.CourseStatistics.UsersProgress))
	for _, u := range response.CourseStatistics.UsersProgress {
		users = append(users, dto.UserCourseProgressResponse{
			UserId:           u.UserId,
			CompletedLessons: int(u.CompletedLessons),
			TotalLessons:     int(u.TotalLessons),
			ProgressPercent:  u.ProgressPercent,
			Completed:        u.Completed,
		})
	}

	return c.JSON(http.StatusOK, dto.CourseStatisticsResponse{
		CourseId: courseId,
		Users:    users,
	})
}

// ReadModuleProgress
// @Summary Получить прогресс модуля
// @Tags Module Progress
// @Produce json
// @Param moduleId path string true "Module ID"
// @Success 200 {object} dto.ModuleProgressResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /progresses/modules/{moduleId} [get]
func (h *LessonProgressHandler) ReadModuleProgress(c echo.Context) error {
	moduleId := c.Param("moduleId")
	userId := c.Get("user_id").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.lessonProgressClient.ReadModuleProgressByUserId(ctx, userId, moduleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.ModuleProgressResponse{
		ModuleId:           response.ModuleProgress.ModuleId,
		TotalLessons:       int(response.ModuleProgress.TotalLessons),
		CompletedLessons:   int(response.ModuleProgress.CompletedLessons),
		ProgressPercent:    response.ModuleProgress.ProgressPercent,
		CompletedLessonIds: response.ModuleProgress.CompletedLessonIds,
	})
}

// ReadModuleStatistics
// @Summary Получить статистику модуля
// @Tags Module Progress
// @Produce json
// @Param moduleId path string true "Module ID"
// @Success 200 {object} dto.ModuleStatisticsResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /progresses/modules/{moduleId}/statistics [get]
func (h *LessonProgressHandler) ReadModuleStatistics(c echo.Context) error {
	moduleId := c.Param("moduleId")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.lessonProgressClient.ReadModuleStatistics(ctx, moduleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	users := make([]dto.UserModuleProgressResponse, 0, len(response.ModuleStatistics.UsersProgress))
	for _, u := range response.ModuleStatistics.UsersProgress {
		users = append(users, dto.UserModuleProgressResponse{
			UserId:           u.UserId,
			CompletedLessons: int(u.CompletedLessons),
			TotalLessons:     int(u.TotalLessons),
			ProgressPercent:  u.ProgressPercent,
			Completed:        u.Completed,
		})
	}

	return c.JSON(http.StatusOK, dto.ModuleStatisticsResponse{
		ModuleId: moduleId,
		Users:    users,
	})
}
