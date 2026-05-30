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

type CompletedTheoryCourseHandler struct {
	completedTheoryCourseClient *course_service.CourseClient
}

func NewCompletedTheoryCourseHandler(client *course_service.CourseClient) *CompletedTheoryCourseHandler {
	return &CompletedTheoryCourseHandler{completedTheoryCourseClient: client}
}

// CreateCompletedTheoryCourse
// @Summary Завершить теоретический курс
// @Tags Completed Theory Courses
// @Accept json
// @Produce json
// @Param request body dto.CompletedTheoryCourseRequest true "Course ID"
// @Success 200 {object} dto.CompletedTheoryCourseResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /completed-theory-courses [post]
func (h *CompletedTheoryCourseHandler) CreateCompletedTheoryCourse(c echo.Context) error {
	var request dto.CompletedTheoryCourseRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid request",
		})
	}

	userId := c.Get("user_id").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.completedTheoryCourseClient.CreateCompletedTheoryCourse(ctx, userId, request.CourseId)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.AlreadyExists:
				return c.JSON(http.StatusConflict, dto.ErrorResponse{
					Error: "completed course already exists",
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

	return c.JSON(http.StatusOK, dto.CompletedTheoryCourseResponse{
		UserId:   response.CompletedTheoryCourse.UserId,
		CourseId: response.CompletedTheoryCourse.CourseId,
	})
}

// ReadAllCompletedTheoryCoursesByUserId
// @Summary Получить завершённые теоретические курсы пользователя
// @Tags Completed Theory Courses
// @Produce json
// @Success 200 {object} dto.CompletedTheoryCourseListResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /completed-theory-courses/my [get]
func (h *CompletedTheoryCourseHandler) ReadAllCompletedTheoryCoursesByUserId(c echo.Context) error {
	userId := c.Get("user_id").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.completedTheoryCourseClient.ReadAllCompletedTheoryCoursesByUserId(ctx, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	completedTheoryCourses := make([]dto.CompletedTheoryCourseResponse, 0, len(response.CompletedTheoryCourses))
	for _, ctc := range response.CompletedTheoryCourses {
		completedTheoryCourses = append(completedTheoryCourses, dto.CompletedTheoryCourseResponse{
			UserId:   ctc.UserId,
			CourseId: ctc.CourseId,
		})
	}

	return c.JSON(http.StatusOK, dto.CompletedTheoryCourseListResponse{
		CompletedTheoryCourses: completedTheoryCourses,
	})
}

// ReadAllCompletedTheoryCoursesByCourseId
// @Summary Получить пользователей завершивших теоретический курс
// @Tags Completed Theory Courses
// @Produce json
// @Param courseId path string true "Course ID"
// @Success 200 {object} dto.CompletedTheoryCourseListResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /completed-theory-courses/courses/{courseId} [get]
func (h *CompletedTheoryCourseHandler) ReadAllCompletedTheoryCoursesByCourseId(c echo.Context) error {
	courseId := c.Param("courseId")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.completedTheoryCourseClient.ReadAllCompletedTheoryCoursesByCourseId(ctx, courseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	completedTheoryCourses := make([]dto.CompletedTheoryCourseResponse, 0, len(response.CompletedTheoryCourses))
	for _, ctc := range response.CompletedTheoryCourses {
		completedTheoryCourses = append(completedTheoryCourses, dto.CompletedTheoryCourseResponse{
			UserId:   ctc.UserId,
			CourseId: ctc.CourseId,
		})
	}

	return c.JSON(http.StatusOK, dto.CompletedTheoryCourseListResponse{
		CompletedTheoryCourses: completedTheoryCourses,
	})
}

func (h *CompletedTheoryCourseHandler) ReadCompletedTheoryCourseByUserIdAndCourseId(c echo.Context) error {
	var request dto.CompletedTheoryCourseRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.completedTheoryCourseClient.ReadCompletedTheoryCourseByUserIdAndCourseId(ctx, request.UserId, request.CourseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.CompletedTheoryCourseResponse{
		UserId:   response.CompletedTheoryCourse.UserId,
		CourseId: response.CompletedTheoryCourse.CourseId,
	})
}
