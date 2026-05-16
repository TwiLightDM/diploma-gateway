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

type CompletedModuleHandler struct {
	completedModuleClient *course_service.CourseClient
}

func NewCompletedModuleHandler(client *course_service.CourseClient) *CompletedModuleHandler {
	return &CompletedModuleHandler{completedModuleClient: client}
}

func (h *CompletedModuleHandler) CreateCompletedModule(c echo.Context) error {
	var request dto.CompletedModuleRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid request",
		})
	}

	userId := c.Get("user_id").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.completedModuleClient.CreateCompletedModule(ctx, userId, request.ModuleId)
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

	return c.JSON(http.StatusOK, dto.CompletedModuleResponse{
		UserId:   response.CompletedModule.UserId,
		ModuleId: response.CompletedModule.ModuleId,
	})
}

func (h *CompletedModuleHandler) ReadCompletedModuleByUserIdAndModuleId(c echo.Context) error {
	var request dto.CompletedModuleRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.completedModuleClient.ReadCompletedModuleByUserIdAndModuleId(ctx, request.UserId, request.ModuleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dto.CompletedModuleResponse{
		UserId:   response.CompletedModule.UserId,
		ModuleId: response.CompletedModule.ModuleId,
	})
}

func (h *CompletedModuleHandler) ReadAllCompletedModulesByUserId(c echo.Context) error {
	userId := c.Get("user_id").(string)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.completedModuleClient.ReadAllCompletedModulesByUserId(ctx, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	completedModules := make([]dto.CompletedModuleResponse, 0, len(response.CompletedModules))
	for _, cm := range response.CompletedModules {
		completedModules = append(completedModules, dto.CompletedModuleResponse{
			UserId:   cm.UserId,
			ModuleId: cm.ModuleId,
		})
	}

	return c.JSON(http.StatusOK, dto.CompletedModuleListResponse{
		CompletedModule: completedModules,
	})
}

func (h *CompletedModuleHandler) ReadAllCompletedModulesByModuleId(c echo.Context) error {
	courseId := c.Param("courseId")

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.completedModuleClient.ReadAllCompletedModulesByModuleId(ctx, courseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
		})
	}

	completedModules := make([]dto.CompletedModuleResponse, 0, len(response.CompletedModules))
	for _, cm := range response.CompletedModules {
		completedModules = append(completedModules, dto.CompletedModuleResponse{
			UserId:   cm.UserId,
			ModuleId: cm.ModuleId,
		})
	}

	return c.JSON(http.StatusOK, dto.CompletedModuleListResponse{
		CompletedModule: completedModules,
	})
}
