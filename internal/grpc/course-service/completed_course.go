package course_service

import (
	"context"

	"github.com/TwiLightDM/diploma-course-service/proto/completedcourseservicepb"
)

func (c *CourseClient) CreateCompletedCourse(ctx context.Context, userId, courseId string) (*completedcourseservicepb.CreateCompletedCourseResponse, error) {
	return c.completedCourse.CreateCompletedCourse(ctx, &completedcourseservicepb.CreateCompletedCourseRequest{
		UserId:   userId,
		CourseId: courseId,
	})
}

func (c *CourseClient) ReadCompletedCourseByUserIdAndCourseId(ctx context.Context, userId, courseId string) (*completedcourseservicepb.ReadCompletedCourseByUserIdAndCourseIdResponse, error) {
	return c.completedCourse.ReadCompletedCourseByUserIdAndCourseId(ctx, &completedcourseservicepb.ReadCompletedCourseByUserIdAndCourseIdRequest{
		UserId:   userId,
		CourseId: courseId,
	})
}

func (c *CourseClient) ReadAllCompletedCoursesByUserId(ctx context.Context, userId string) (*completedcourseservicepb.ReadAllCompletedCoursesByUserIdResponse, error) {
	return c.completedCourse.ReadAllCompletedCoursesByUserId(ctx, &completedcourseservicepb.ReadAllCompletedCoursesByUserIdRequest{
		UserId: userId,
	})
}

func (c *CourseClient) ReadAllCompletedCoursesByCourseId(ctx context.Context, courseId string) (*completedcourseservicepb.ReadAllCompletedCoursesByCourseIdResponse, error) {
	return c.completedCourse.ReadAllCompletedCoursesByCourseId(ctx, &completedcourseservicepb.ReadAllCompletedCoursesByCourseIdRequest{
		CourseId: courseId,
	})
}
