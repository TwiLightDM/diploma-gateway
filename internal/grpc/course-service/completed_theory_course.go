package course_service

import (
	"context"

	"github.com/TwiLightDM/diploma-course-service/proto/completedtheorycourseservicepb"
)

func (c *CourseClient) CreateCompletedTheoryCourse(ctx context.Context, userId, courseId string) (*completedtheorycourseservicepb.CreateCompletedTheoryCourseResponse, error) {
	return c.completedTheoryCourse.CreateCompletedTheoryCourse(ctx, &completedtheorycourseservicepb.CreateCompletedTheoryCourseRequest{
		UserId:   userId,
		CourseId: courseId,
	})
}

func (c *CourseClient) ReadCompletedTheoryCourseByUserIdAndCourseId(ctx context.Context, userId, courseId string) (*completedtheorycourseservicepb.ReadCompletedTheoryCourseByUserIdAndCourseIdResponse, error) {
	return c.completedTheoryCourse.ReadCompletedTheoryCourseByUserIdAndCourseId(ctx, &completedtheorycourseservicepb.ReadCompletedTheoryCourseByUserIdAndCourseIdRequest{
		UserId:   userId,
		CourseId: courseId,
	})
}

func (c *CourseClient) ReadAllCompletedTheoryCoursesByUserId(ctx context.Context, userId string) (*completedtheorycourseservicepb.ReadAllCompletedTheoryCoursesByUserIdResponse, error) {
	return c.completedTheoryCourse.ReadAllCompletedTheoryCoursesByUserId(ctx, &completedtheorycourseservicepb.ReadAllCompletedTheoryCoursesByUserIdRequest{
		UserId: userId,
	})
}

func (c *CourseClient) ReadAllCompletedTheoryCoursesByCourseId(ctx context.Context, courseId string) (*completedtheorycourseservicepb.ReadAllCompletedTheoryCoursesByCourseIdResponse, error) {
	return c.completedTheoryCourse.ReadAllCompletedTheoryCoursesByCourseId(ctx, &completedtheorycourseservicepb.ReadAllCompletedTheoryCoursesByCourseIdRequest{
		CourseId: courseId,
	})
}
