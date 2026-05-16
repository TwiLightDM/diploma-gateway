package course_service

import (
	"context"

	"github.com/TwiLightDM/diploma-course-service/proto/completedmoduleservicepb"
)

func (c *CourseClient) CreateCompletedModule(ctx context.Context, userId, courseId string) (*completedmoduleservicepb.CreateCompletedModuleResponse, error) {
	return c.completedModule.CreateCompletedModule(ctx, &completedmoduleservicepb.CreateCompletedModuleRequest{
		UserId:   userId,
		ModuleId: courseId,
	})
}

func (c *CourseClient) ReadCompletedModuleByUserIdAndModuleId(ctx context.Context, userId, moduleId string) (*completedmoduleservicepb.ReadCompletedModuleByUserIdAndModuleIdResponse, error) {
	return c.completedModule.ReadCompletedModuleByUserIdAndModuleId(ctx, &completedmoduleservicepb.ReadCompletedModuleByUserIdAndModuleIdRequest{
		UserId:   userId,
		ModuleId: moduleId,
	})
}

func (c *CourseClient) ReadAllCompletedModulesByUserId(ctx context.Context, userId string) (*completedmoduleservicepb.ReadAllCompletedModulesByUserIdResponse, error) {
	return c.completedModule.ReadAllCompletedModulesByUserId(ctx, &completedmoduleservicepb.ReadAllCompletedModulesByUserIdRequest{
		UserId: userId,
	})
}

func (c *CourseClient) ReadAllCompletedModulesByModuleId(ctx context.Context, moduleId string) (*completedmoduleservicepb.ReadAllCompletedModulesByModuleIdResponse, error) {
	return c.completedModule.ReadAllCompletedModulesByModuleId(ctx, &completedmoduleservicepb.ReadAllCompletedModulesByModuleIdRequest{
		ModuleId: moduleId,
	})
}
