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

type CompletedCourseHandler struct {
	completedCourseClient *course_service.CourseClient
}

func NewCompletedCourseHandler(client *course_service.CourseClient) *CompletedCourseHandler {
	return &CompletedCourseHandler{completedCourseClient: client}
}

func (h *CompletedCourseHandler) CreateCompletedCourse(c echo.Context) error {
	var request dto.CompletedCourseRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid request",
		})
	}

	userId := c.Get("user_id").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.completedCourseClient.CreateCompletedCourse(ctx, userId, request.CourseId)
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

	return c.JSON(http.StatusOK, dto.CompletedCourseResponse{
		UserId:   response.CompletedCourse.UserId,
		CourseId: response.CompletedCourse.CourseId,
	})
}

func (h *CompletedCourseHandler) ReadCompletedCourseByUserIdAndCourseId(c echo.Context) error {
	var request dto.CompletedCourseRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.completedCourseClient.ReadCompletedCourseByUserIdAndCourseId(ctx, request.UserId, request.CourseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.CompletedCourseResponse{
		UserId:   response.CompletedCourse.UserId,
		CourseId: response.CompletedCourse.CourseId,
	})
}

func (h *CompletedCourseHandler) ReadAllCompletedCoursesByUserId(c echo.Context) error {
	userId := c.Get("user_id").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.completedCourseClient.ReadAllCompletedCoursesByUserId(ctx, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	completedCourses := make([]dto.CompletedCourseResponse, 0, len(response.CompletedCourses))
	for _, cc := range response.CompletedCourses {
		completedCourses = append(completedCourses, dto.CompletedCourseResponse{
			UserId:   cc.UserId,
			CourseId: cc.CourseId,
		})
	}

	return c.JSON(http.StatusOK, dto.CompletedCourseListResponse{
		CompletedCourses: completedCourses,
	})
}

func (h *CompletedCourseHandler) ReadAllCompletedCoursesByCourseId(c echo.Context) error {
	courseId := c.Param("courseId")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.completedCourseClient.ReadAllCompletedCoursesByCourseId(ctx, courseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	completedCourses := make([]dto.CompletedCourseResponse, 0, len(response.CompletedCourses))
	for _, cc := range response.CompletedCourses {
		completedCourses = append(completedCourses, dto.CompletedCourseResponse{
			UserId:   cc.UserId,
			CourseId: cc.CourseId,
		})
	}

	return c.JSON(http.StatusOK, dto.CompletedCourseListResponse{
		CompletedCourses: completedCourses,
	})
}
