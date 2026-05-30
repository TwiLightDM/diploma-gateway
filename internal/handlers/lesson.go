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

type LessonHandler struct {
	courseClient *course_service.CourseClient
}

func NewLessonHandler(courseClient *course_service.CourseClient) *LessonHandler {
	return &LessonHandler{courseClient: courseClient}
}

// CreateLesson
// @Summary Создать урок
// @Tags Lessons
// @Accept json
// @Produce json
// @Param request body dto.LessonRequest true "Данные урока"
// @Success 200 {object} dto.LessonResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /lessons [post]
func (h *LessonHandler) CreateLesson(c echo.Context) error {
	var request dto.LessonRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.courseClient.CreateLesson(ctx, request.Title, request.Description, request.Content, request.ModuleId)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.AlreadyExists:
				return c.JSON(http.StatusConflict, dto.ErrorResponse{Error: "course already exists"})
			default:
				return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			}
		}
	}

	return c.JSON(http.StatusOK, dto.LessonResponse{
		Id:          response.Lesson.Id,
		Title:       response.Lesson.Title,
		Description: response.Lesson.Description,
		Content:     response.Lesson.Content,
		Position:    response.Lesson.Position,
		ModuleId:    response.Lesson.ModuleId,
	})
}

// ReadLesson
// @Summary Получить урок
// @Tags Lessons
// @Produce json
// @Param id path string true "Lesson ID"
// @Success 200 {object} dto.LessonResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /lessons/{id} [get]
func (h *LessonHandler) ReadLesson(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadLesson(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.LessonResponse{
		Id:          response.Lesson.Id,
		Title:       response.Lesson.Title,
		Description: response.Lesson.Description,
		Content:     response.Lesson.Content,
		Position:    response.Lesson.Position,
		ModuleId:    response.Lesson.ModuleId,
	})
}

// ReadAllLessonsByCourseId
// @Summary Получить уроки модуля
// @Tags Lessons
// @Produce json
// @Param module_id path string true "Module ID"
// @Success 200 {object} dto.LessonListResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /lessons/modules/{module_id} [get]
func (h *LessonHandler) ReadAllLessonsByCourseId(c echo.Context) error {
	moduleId := c.Param("module_id")
	if moduleId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadAllLessonsByModuleId(ctx, moduleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	lessons := make([]dto.LessonResponse, 0, len(response.Lessons))
	for _, lesson := range response.Lessons {
		lessons = append(lessons, dto.LessonResponse{
			Id:          lesson.Id,
			Title:       lesson.Title,
			Description: lesson.Description,
			Position:    lesson.Position,
			ModuleId:    lesson.ModuleId,
		})
	}

	return c.JSON(http.StatusOK, dto.LessonListResponse{
		Lessons: lessons,
	})
}

// UpdateLesson
// @Summary Обновить урок
// @Tags Lessons
// @Accept json
// @Produce json
// @Param id path string true "Lesson ID"
// @Param request body dto.LessonRequest true "Новые данные"
// @Success 200 {object} dto.LessonResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /lessons/{id} [patch]
func (h *LessonHandler) UpdateLesson(c echo.Context) error {
	var request dto.LessonRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.courseClient.UpdateLesson(ctx, id, request.Title, request.Description, request.Content, request.Position)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.LessonResponse{
		Id:          response.Lesson.Id,
		Title:       response.Lesson.Title,
		Description: response.Lesson.Description,
		Content:     response.Lesson.Content,
		Position:    response.Lesson.Position,
		ModuleId:    response.Lesson.ModuleId,
	})
}

// DeleteLesson
// @Summary Удалить урок
// @Tags Lessons
// @Produce json
// @Param id path string true "Lesson ID"
// @Success 204
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /lessons/{id} [delete]
func (h *LessonHandler) DeleteLesson(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	_, err := h.courseClient.DeleteLesson(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}
