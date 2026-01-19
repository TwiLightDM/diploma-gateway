package handlers

import (
	"context"
	"github.com/TwiLightDM/diploma-gateway/internal/dto"
	"github.com/TwiLightDM/diploma-gateway/internal/grpc/course-service"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type ModuleHandler struct {
	courseClient *course_service.CourseClient
}

func NewModuleHandler(courseClient *course_service.CourseClient) *ModuleHandler {
	return &ModuleHandler{courseClient: courseClient}
}

func (h *ModuleHandler) CreateModule(c echo.Context) error {
	var request dto.ModuleRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.SignUpResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := h.courseClient.CreateModule(ctx, request.Title, request.Description, request.CourseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ModuleResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.ModuleResponse{
		Id:          response.Module.Id,
		Title:       response.Module.Title,
		Description: response.Module.Description,
		Position:    response.Module.Position,
		CourseId:    response.Module.CourseId,
	})
}

func (h *ModuleHandler) ReadModule(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ModuleResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadModule(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ModuleResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.ModuleResponse{
		Id:          response.Module.Id,
		Title:       response.Module.Title,
		Description: response.Module.Description,
		Position:    response.Module.Position,
		CourseId:    response.Module.CourseId,
	})
}

func (h *ModuleHandler) ReadAllModulesByCourseId(c echo.Context) error {
	ownerId := c.Param("id")
	if ownerId == "" {
		return c.JSON(http.StatusBadRequest, dto.ModuleResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadAllModulesByCourseId(ctx, ownerId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ModuleResponse{Error: err.Error()})
	}

	modules := make([]dto.ModuleResponse, 0, len(response.Modules))
	for _, module := range response.Modules {
		modules = append(modules, dto.ModuleResponse{
			Id:          module.Id,
			Title:       module.Title,
			Description: module.Description,
			Position:    module.Position,
			CourseId:    module.CourseId,
		})
	}

	return c.JSON(http.StatusOK, dto.ModuleListResponse{
		Modules: modules,
	})
}

func (h *ModuleHandler) UpdateModule(c echo.Context) error {
	var request dto.ModuleRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ModuleResponse{Error: "invalid request"})
	}

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ModuleResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := h.courseClient.UpdateModule(ctx, id, request.Title, request.Description, request.Position)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ModuleResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.ModuleResponse{
		Id:          response.Module.Id,
		Title:       response.Module.Title,
		Description: response.Module.Description,
		Position:    response.Module.Position,
		CourseId:    response.Module.CourseId,
	})
}

func (h *ModuleHandler) DeleteModule(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ModuleResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := h.courseClient.DeleteModule(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ModuleResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusNoContent, nil)
}
