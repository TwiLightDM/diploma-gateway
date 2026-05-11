package course_service

import (
	"context"

	"github.com/TwiLightDM/diploma-course-service/proto/taskattemptservicepb"
)

func (c *CourseClient) SubmitTaskAttempt(
	ctx context.Context,
	userId,
	courseId,
	moduleId string,
	answers []*taskattemptservicepb.TaskAttemptAnswer,
) (*taskattemptservicepb.SubmitTaskAttemptResponse, error) {
	return c.taskAttempt.SubmitTaskAttempt(ctx, &taskattemptservicepb.SubmitTaskAttemptRequest{
		UserId:   userId,
		CourseId: courseId,
		ModuleId: moduleId,
		Answers:  answers,
	})
}

func (c *CourseClient) ReadTaskAttempt(ctx context.Context, id string) (*taskattemptservicepb.ReadTaskAttemptResponse, error) {
	return c.taskAttempt.ReadTaskAttempt(ctx, &taskattemptservicepb.ReadTaskAttemptRequest{
		Id: id,
	})
}

func (c *CourseClient) ReadAllTaskAttemptsByUserIdAndModuleId(ctx context.Context, userId, moduleId string) (*taskattemptservicepb.ReadAllTaskAttemptsByUserIdAndModuleIdResponse, error) {
	return c.taskAttempt.ReadAllTaskAttemptsByUserIdAndModuleId(ctx, &taskattemptservicepb.ReadAllTaskAttemptsByUserIdAndModuleIdRequest{
		UserId:   userId,
		ModuleId: moduleId,
	})
}

func (c *CourseClient) ReadAllTaskAttemptsByUserIdAndCourseId(ctx context.Context, userId, courseId string) (*taskattemptservicepb.ReadAllTaskAttemptsByUserIdAndCourseIdResponse, error) {
	return c.taskAttempt.ReadAllTaskAttemptsByUserIdAndCourseId(ctx, &taskattemptservicepb.ReadAllTaskAttemptsByUserIdAndCourseIdRequest{
		UserId:   userId,
		CourseId: courseId,
	})
}
