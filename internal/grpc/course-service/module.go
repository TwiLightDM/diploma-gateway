package course_service

import (
	"context"
	"github.com/TwiLightDM/diploma-course-service/proto/moduleservicepb"
)

func (c *CourseClient) CreateModule(ctx context.Context, title, description, courseId string) (*moduleservicepb.CreateModuleResponse, error) {
	return c.module.CreateModule(ctx, &moduleservicepb.CreateModuleRequest{
		Title:       title,
		Description: description,
		CourseId:    courseId,
	})
}

func (c *CourseClient) ReadModule(ctx context.Context, id string) (*moduleservicepb.ReadModuleResponse, error) {
	return c.module.ReadModule(ctx, &moduleservicepb.ReadModuleRequest{
		Id: id,
	})
}

func (c *CourseClient) ReadAllModulesByCourseId(ctx context.Context, courseId string) (*moduleservicepb.ReadAllModulesByCourseIdResponse, error) {
	return c.module.ReadAllModulesByCourseId(ctx, &moduleservicepb.ReadAllModulesByCourseIdRequest{
		CourseId: courseId,
	})
}

func (c *CourseClient) UpdateModule(ctx context.Context, id, title, description string, position int64) (*moduleservicepb.UpdateModuleResponse, error) {
	return c.module.UpdateModule(ctx, &moduleservicepb.UpdateModuleRequest{
		Id:          id,
		Title:       title,
		Description: description,
		Position:    position,
	})
}

func (c *CourseClient) DeleteModule(ctx context.Context, id string) (*moduleservicepb.DeleteModuleResponse, error) {
	return c.module.DeleteModule(ctx, &moduleservicepb.DeleteModuleRequest{
		Id: id,
	})
}
