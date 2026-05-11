package course_service

import (
	"context"

	"github.com/TwiLightDM/diploma-course-service/proto/taskservicepb"
	"github.com/TwiLightDM/diploma-gateway/internal/errs"
)

func (c *CourseClient) CreateTask(
	ctx context.Context,
	title, courseId, moduleId, correctTextAnswer, taskType string,
	options []*taskservicepb.TaskOption,
) (*taskservicepb.CreateTaskResponse, error) {
	var tType taskservicepb.TaskType

	switch taskType {
	case "TEXT_INPUT":
		tType = taskservicepb.TaskType_TASK_TYPE_TEXT_INPUT
	case "CHOICE":
		tType = taskservicepb.TaskType_TASK_TYPE_CHOICE
	default:
		return nil, errs.ErrInvalidTaskType
	}

	return c.task.CreateTask(ctx, &taskservicepb.CreateTaskRequest{
		Title:             title,
		CourseId:          courseId,
		ModuleId:          moduleId,
		Type:              tType,
		CorrectTextAnswer: correctTextAnswer,
		Options:           options,
	})
}

func (c *CourseClient) ReadTask(ctx context.Context, id string) (*taskservicepb.ReadTaskResponse, error) {
	return c.task.ReadTask(ctx, &taskservicepb.ReadTaskRequest{
		Id: id,
	})
}

func (c *CourseClient) ReadAllTasksByCourseId(ctx context.Context, courseId string) (*taskservicepb.ReadAllTasksByCourseIdResponse, error) {
	return c.task.ReadAllTasksByCourseId(ctx, &taskservicepb.ReadAllTasksByCourseIdRequest{
		CourseId: courseId,
	})
}

func (c *CourseClient) ReadAllTasksByModuleId(ctx context.Context, moduleId string) (*taskservicepb.ReadAllTasksByModuleIdResponse, error) {
	return c.task.ReadAllTasksByModuleId(ctx, &taskservicepb.ReadAllTasksByModuleIdRequest{
		ModuleId: moduleId,
	})
}

func (c *CourseClient) UpdateTask(
	ctx context.Context,
	id, title, correctTextAnswer string,
	options []*taskservicepb.TaskOption,
) (*taskservicepb.UpdateTaskResponse, error) {
	return c.task.UpdateTask(ctx, &taskservicepb.UpdateTaskRequest{
		Id:                id,
		Title:             title,
		CorrectTextAnswer: correctTextAnswer,
		Options:           options,
	})
}

func (c *CourseClient) DeleteTask(ctx context.Context, id string) (*taskservicepb.DeleteTaskResponse, error) {
	return c.task.DeleteTask(ctx, &taskservicepb.DeleteTaskRequest{
		Id: id,
	})
}
