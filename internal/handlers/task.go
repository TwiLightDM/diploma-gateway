package handlers

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/TwiLightDM/diploma-course-service/proto/taskservicepb"
	"github.com/TwiLightDM/diploma-gateway/internal/dto"
	"github.com/TwiLightDM/diploma-gateway/internal/errs"
	"github.com/TwiLightDM/diploma-gateway/internal/grpc/course-service"
	"github.com/labstack/echo/v4"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskHandler struct {
	courseClient *course_service.CourseClient
}

func NewTaskHandler(courseClient *course_service.CourseClient) *TaskHandler {
	return &TaskHandler{courseClient: courseClient}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var request dto.TaskRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := make([]*taskservicepb.TaskOption, 0, len(request.Options))
	for _, option := range request.Options {
		options = append(options, &taskservicepb.TaskOption{
			Text:      option.Text,
			IsCorrect: option.IsCorrect,
		})
	}

	response, err := h.courseClient.CreateTask(
		ctx,
		request.Title,
		request.CourseId,
		request.ModuleId,
		request.CorrectTextAnswer,
		request.Type,
		options,
	)
	if err != nil {
		if errors.Is(err, errs.ErrInvalidTaskType) {
			return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		}

		if st, ok := status.FromError(err); ok {
			switch st.Code() {

			case codes.AlreadyExists:
				return c.JSON(http.StatusConflict, dto.ErrorResponse{Error: "task already exists"})

			default:
				return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
			}
		}
	}

	dtoOptions := make([]dto.TaskOptionResponse, 0, len(response.Task.Options))
	for _, option := range response.Task.Options {
		dtoOptions = append(dtoOptions, dto.TaskOptionResponse{
			Id:        option.Id,
			Text:      option.Text,
			IsCorrect: option.IsCorrect,
		})
	}

	return c.JSON(http.StatusOK, dto.TaskResponse{
		Id:                response.Task.Id,
		Title:             response.Task.Title,
		CourseId:          response.Task.CourseId,
		ModuleId:          response.Task.ModuleId,
		Type:              response.Task.Type.String(),
		CorrectTextAnswer: response.Task.CorrectTextAnswer,
		Options:           dtoOptions,
	},
	)
}

func (h *TaskHandler) ReadTask(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadTask(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	dtoOptions := make([]dto.TaskOptionResponse, 0, len(response.Task.Options))
	for _, option := range response.Task.Options {
		dtoOptions = append(dtoOptions, dto.TaskOptionResponse{
			Id:        option.Id,
			Text:      option.Text,
			IsCorrect: option.IsCorrect,
		})
	}

	return c.JSON(http.StatusOK, dto.TaskResponse{
		Id:                response.Task.Id,
		Title:             response.Task.Title,
		CourseId:          response.Task.CourseId,
		ModuleId:          response.Task.ModuleId,
		Type:              response.Task.Type.String(),
		CorrectTextAnswer: response.Task.CorrectTextAnswer,
		Options:           dtoOptions,
	},
	)
}

func (h *TaskHandler) ReadTasksByModuleId(c echo.Context) error {
	moduleId := c.Param("module_id")
	if moduleId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadAllTasksByModuleId(ctx, moduleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	tasks := make([]dto.TaskResponse, 0, len(response.Tasks))

	for _, task := range response.Tasks {
		dtoOptions := make([]dto.TaskOptionResponse, 0, len(task.Options))
		for _, option := range task.Options {
			dtoOptions = append(dtoOptions, dto.TaskOptionResponse{
				Id:        option.Id,
				Text:      option.Text,
				IsCorrect: option.IsCorrect,
			})
		}

		tasks = append(tasks, dto.TaskResponse{
			Id:                task.Id,
			Title:             task.Title,
			CourseId:          task.CourseId,
			ModuleId:          task.ModuleId,
			Type:              task.Type.String(),
			CorrectTextAnswer: task.CorrectTextAnswer,
			Options:           dtoOptions,
		})
	}

	return c.JSON(http.StatusOK, dto.TaskListResponse{
		Tasks: tasks,
	},
	)
}

func (h *TaskHandler) ReadTasksByCourseId(c echo.Context) error {
	courseId := c.Param("course_id")
	if courseId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadAllTasksByCourseId(
		ctx,
		courseId,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	tasks := make([]dto.TaskResponse, 0, len(response.Tasks))

	for _, task := range response.Tasks {
		dtoOptions := make([]dto.TaskOptionResponse, 0, len(task.Options))
		for _, option := range task.Options {
			dtoOptions = append(dtoOptions, dto.TaskOptionResponse{
				Id:        option.Id,
				Text:      option.Text,
				IsCorrect: option.IsCorrect,
			})
		}
		tasks = append(tasks, dto.TaskResponse{
			Id:                task.Id,
			Title:             task.Title,
			CourseId:          task.CourseId,
			ModuleId:          task.ModuleId,
			Type:              task.Type.String(),
			CorrectTextAnswer: task.CorrectTextAnswer,
			Options:           dtoOptions,
		})
	}

	return c.JSON(http.StatusOK, dto.TaskListResponse{
		Tasks: tasks,
	},
	)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	var request dto.TaskRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := make([]*taskservicepb.TaskOption, 0, len(request.Options))
	for _, option := range request.Options {
		options = append(options, &taskservicepb.TaskOption{
			Text:      option.Text,
			IsCorrect: option.IsCorrect,
		})
	}

	response, err := h.courseClient.UpdateTask(
		ctx,
		id,
		request.Title,
		request.CorrectTextAnswer,
		options,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	dtoOptions := make([]dto.TaskOptionResponse, 0, len(response.Task.Options))
	for _, option := range response.Task.Options {
		dtoOptions = append(dtoOptions, dto.TaskOptionResponse{
			Id:        option.Id,
			Text:      option.Text,
			IsCorrect: option.IsCorrect,
		})
	}

	return c.JSON(http.StatusOK, dto.TaskResponse{
		Id:                response.Task.Id,
		Title:             response.Task.Title,
		CourseId:          response.Task.CourseId,
		ModuleId:          response.Task.ModuleId,
		Type:              response.Task.Type.String(),
		CorrectTextAnswer: response.Task.CorrectTextAnswer,
		Options:           dtoOptions,
	},
	)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := h.courseClient.DeleteTask(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
