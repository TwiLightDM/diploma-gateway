package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/TwiLightDM/diploma-course-service/proto/taskattemptservicepb"
	"github.com/TwiLightDM/diploma-gateway/internal/dto"
	"github.com/TwiLightDM/diploma-gateway/internal/grpc/course-service"

	"github.com/labstack/echo/v4"
)

type TaskAttemptHandler struct {
	courseClient *course_service.CourseClient
}

func NewTaskAttemptHandler(courseClient *course_service.CourseClient) *TaskAttemptHandler {
	return &TaskAttemptHandler{courseClient: courseClient}
}

func (h *TaskAttemptHandler) SubmitTaskAttempt(c echo.Context) error {
	var request dto.TaskAttemptRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	answers := make([]*taskattemptservicepb.TaskAttemptAnswer, 0, len(request.Answers))
	for _, answer := range request.Answers {
		answers = append(answers, &taskattemptservicepb.TaskAttemptAnswer{
			TaskId:            answer.TaskId,
			TextAnswer:        answer.TextAnswer,
			SelectedOptionIds: answer.SelectedOptionIds,
		},
		)
	}

	response, err := h.courseClient.SubmitTaskAttempt(
		ctx,
		request.UserId,
		request.CourseId,
		request.ModuleId,
		answers,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, mapTaskAttemptResponse(response.TaskAttempt))
}

func (h *TaskAttemptHandler) ReadTaskAttempt(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadTaskAttempt(ctx, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, mapTaskAttemptResponse(response.TaskAttempt))
}

func (h *TaskAttemptHandler) ReadAllTaskAttemptsByUserIdAndModuleId(c echo.Context) error {
	userId := c.Param("user_id")
	if userId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	moduleId := c.Param("module_id")
	if moduleId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadAllTaskAttemptsByUserIdAndModuleId(ctx, userId, moduleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	taskAttempts := make([]dto.TaskAttemptResponse, 0, len(response.TaskAttempts))
	for _, taskAttempt := range response.TaskAttempts {
		taskAttempts = append(taskAttempts, mapTaskAttemptResponse(taskAttempt))
	}

	return c.JSON(http.StatusOK, dto.TaskAttemptListResponse{
		TaskAttempts: taskAttempts,
	},
	)
}

func (h *TaskAttemptHandler) ReadAllTaskAttemptsByUserIdAndCourseId(c echo.Context) error {
	userId := c.Param("user_id")
	if userId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	courseId := c.Param("course_id")
	if courseId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadAllTaskAttemptsByUserIdAndCourseId(ctx, userId, courseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	taskAttempts := make([]dto.TaskAttemptResponse, 0, len(response.TaskAttempts))
	for _, taskAttempt := range response.TaskAttempts {
		taskAttempts = append(taskAttempts, mapTaskAttemptResponse(taskAttempt))
	}

	return c.JSON(http.StatusOK, dto.TaskAttemptListResponse{
		TaskAttempts: taskAttempts,
	},
	)
}

func (h *TaskAttemptHandler) ReadAllTaskAttemptsByYourIdAndModuleId(c echo.Context) error {
	userId := c.Get("user_id").(string)
	if userId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	moduleId := c.Param("module_id")
	if moduleId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadAllTaskAttemptsByUserIdAndModuleId(ctx, userId, moduleId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	taskAttempts := make([]dto.TaskAttemptResponse, 0, len(response.TaskAttempts))
	for _, taskAttempt := range response.TaskAttempts {
		taskAttempts = append(taskAttempts, mapTaskAttemptResponse(taskAttempt))
	}

	return c.JSON(http.StatusOK, dto.TaskAttemptListResponse{
		TaskAttempts: taskAttempts,
	},
	)
}

func (h *TaskAttemptHandler) ReadAllTaskAttemptsByYourIdAndCourseId(c echo.Context) error {
	userId := c.Get("user_id").(string)
	if userId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}
	fmt.Println(userId)
	courseId := c.Param("course_id")
	if courseId == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "invalid request"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	response, err := h.courseClient.ReadAllTaskAttemptsByUserIdAndCourseId(ctx, userId, courseId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
	}

	taskAttempts := make([]dto.TaskAttemptResponse, 0, len(response.TaskAttempts))
	for _, taskAttempt := range response.TaskAttempts {
		taskAttempts = append(taskAttempts, mapTaskAttemptResponse(taskAttempt))
	}

	return c.JSON(http.StatusOK, dto.TaskAttemptListResponse{
		TaskAttempts: taskAttempts,
	},
	)
}

func mapTaskAttemptResponse(taskAttempt *taskattemptservicepb.TaskAttempt) dto.TaskAttemptResponse {
	answers := make([]dto.TaskAttemptAnswerResponse, 0, len(taskAttempt.Answers))
	for _, answer := range taskAttempt.Answers {
		answers = append(
			answers, dto.TaskAttemptAnswerResponse{
				TaskId:            answer.TaskId,
				TextAnswer:        answer.TextAnswer,
				SelectedOptionIds: answer.SelectedOptionIds,
				IsCorrect:         answer.IsCorrect,
			},
		)
	}

	return dto.TaskAttemptResponse{
		Id:             taskAttempt.Id,
		UserId:         taskAttempt.UserId,
		CourseId:       taskAttempt.CourseId,
		ModuleId:       taskAttempt.ModuleId,
		AttemptNumber:  int(taskAttempt.AttemptNumber),
		Answers:        answers,
		CorrectAnswers: int(taskAttempt.CorrectAnswers),
		TotalQuestions: int(taskAttempt.TotalQuestions),
		Score:          taskAttempt.Score,
	}
}
