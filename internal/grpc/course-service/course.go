package course_service

import (
	"context"
	"github.com/TwiLightDM/diploma-course-service/proto/courseservicepb"
)

func (c *CourseClient) CreateCourse(ctx context.Context, title, description, accessType, ownerId string) (*courseservicepb.CreateCourseResponse, error) {
	return c.course.CreateCourse(ctx, &courseservicepb.CreateCourseRequest{
		Title:       title,
		Description: description,
		AccessType:  accessType,
		OwnerId:     ownerId,
	})
}

func (c *CourseClient) ReadCourse(ctx context.Context, id string) (*courseservicepb.ReadCourseResponse, error) {
	return c.course.ReadCourse(ctx, &courseservicepb.ReadCourseRequest{
		Id: id,
	})
}

func (c *CourseClient) ReadAllCourseByOwnerId(ctx context.Context, ownerId string) (*courseservicepb.ReadAllCoursesByOwnerIdResponse, error) {
	return c.course.ReadAllCoursesByOwnerId(ctx, &courseservicepb.ReadAllCoursesByOwnerIdRequest{
		OwnerId: ownerId,
	})
}

func (c *CourseClient) UpdateCourse(ctx context.Context, id, title, description, accessType string) (*courseservicepb.UpdateCourseResponse, error) {
	return c.course.UpdateCourse(ctx, &courseservicepb.UpdateCourseRequest{
		Id:          id,
		Title:       title,
		Description: description,
		AccessType:  accessType,
	})
}

func (c *CourseClient) UpdatePublishedAt(ctx context.Context, id string) (*courseservicepb.UpdateCourseResponse, error) {
	return c.course.UpdatePublishedAt(ctx, &courseservicepb.UpdatePublishedAtRequest{
		Id: id,
	})
}

func (c *CourseClient) DeleteCourse(ctx context.Context, id string) (*courseservicepb.DeleteCourseResponse, error) {
	return c.course.DeleteCourse(ctx, &courseservicepb.DeleteCourseRequest{
		Id: id,
	})
}
