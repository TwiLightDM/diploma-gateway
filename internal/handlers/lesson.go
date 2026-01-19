package handlers

import (
	"context"
	"github.com/TwiLightDM/diploma-gateway/internal/dto"
	"github.com/TwiLightDM/diploma-gateway/internal/grpc/course-service"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type LessonHandler struct {
	courseClient *course_service.CourseClient
}

func NewLessonHandler(courseClient *course_service.CourseClient) *LessonHandler {
	return &LessonHandler{courseClient: courseClient}
}

func (h *LessonHandler) CreateLesson(c echo.Context) error {
	var request dto.LessonRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.SignUpResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := h.courseClient.CreateLesson(ctx, request.Title, request.Description, request.ModuleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.LessonResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.LessonResponse{
		Id:          response.Lesson.Id,
		Title:       response.Lesson.Title,
		Description: response.Lesson.Description,
		Position:    response.Lesson.Position,
		ModuleId:    response.Lesson.ModuleId,
	})
}

func (h *LessonHandler) ReadLesson(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.LessonResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadLesson(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.LessonResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.LessonResponse{
		Id:          response.Lesson.Id,
		Title:       response.Lesson.Title,
		Description: response.Lesson.Description,
		Position:    response.Lesson.Position,
		ModuleId:    response.Lesson.ModuleId,
	})
}

func (h *LessonHandler) ReadAllLessonsByCourseId(c echo.Context) error {
	ownerId := c.Param("id")
	if ownerId == "" {
		return c.JSON(http.StatusBadRequest, dto.LessonResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadAllLessonsByModuleId(ctx, ownerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.LessonResponse{Error: err.Error()})
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

func (h *LessonHandler) UpdateLesson(c echo.Context) error {
	var request dto.LessonRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.LessonResponse{Error: "invalid request"})
	}

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.LessonResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := h.courseClient.UpdateLesson(ctx, id, request.Title, request.Description, request.Position)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.LessonResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.LessonResponse{
		Id:          response.Lesson.Id,
		Title:       response.Lesson.Title,
		Description: response.Lesson.Description,
		Position:    response.Lesson.Position,
		ModuleId:    response.Lesson.ModuleId,
	})
}

func (h *LessonHandler) DeleteLesson(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.LessonResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := h.courseClient.DeleteLesson(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.LessonResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}
