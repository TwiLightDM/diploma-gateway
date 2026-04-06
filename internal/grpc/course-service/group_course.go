package course_service

import (
	"context"

	"github.com/TwiLightDM/diploma-course-service/proto/groupcourseservicepb"
)

func (c *CourseClient) CreateGroupCourse(ctx context.Context, courseId, groupId string) (*groupcourseservicepb.CreateGroupCourseResponse, error) {
	return c.groupCourse.CreateGroupCourse(ctx, &groupcourseservicepb.CreateGroupCourseRequest{
		GroupId:  groupId,
		CourseId: courseId,
	})
}

func (c *CourseClient) ReadAllGroupCoursesByGroupId(ctx context.Context, groupId string) (*groupcourseservicepb.ReadAllGroupCoursesByGroupIdResponse, error) {
	return c.groupCourse.ReadAllGroupCoursesByGroupId(ctx, &groupcourseservicepb.ReadAllGroupCoursesByGroupIdRequest{
		GroupId: groupId,
	})
}

func (c *CourseClient) ReadAllGroupCoursesByCourseId(ctx context.Context, courseId string) (*groupcourseservicepb.ReadAllGroupCoursesByCourseIdResponse, error) {
	return c.groupCourse.ReadAllGroupCoursesByCourseId(ctx, &groupcourseservicepb.ReadAllGroupCoursesByCourseIdRequest{
		CourseId: courseId,
	})
}

func (c *CourseClient) DeleteGroupCourse(ctx context.Context, id string) (*groupcourseservicepb.DeleteGroupCourseResponse, error) {
	return c.groupCourse.DeleteGroupCourse(ctx, &groupcourseservicepb.DeleteGroupCourseRequest{
		Id: id,
	})
}
