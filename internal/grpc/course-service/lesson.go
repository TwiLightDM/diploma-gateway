package course_service

import (
	"context"
	"github.com/TwiLightDM/diploma-course-service/proto/lessonservicepb"
)

func (c *CourseClient) CreateLesson(ctx context.Context, title, description, moduleId string) (*lessonservicepb.CreateLessonResponse, error) {
	return c.lesson.CreateLesson(ctx, &lessonservicepb.CreateLessonRequest{
		Title:       title,
		Description: description,
		ModuleId:    moduleId,
	})
}

func (c *CourseClient) ReadLesson(ctx context.Context, id string) (*lessonservicepb.ReadLessonResponse, error) {
	return c.lesson.ReadLesson(ctx, &lessonservicepb.ReadLessonRequest{
		Id: id,
	})
}

func (c *CourseClient) ReadAllLessonsByModuleId(ctx context.Context, moduleId string) (*lessonservicepb.ReadAllLessonsByModuleIdResponse, error) {
	return c.lesson.ReadAllLessonsByModuleId(ctx, &lessonservicepb.ReadAllLessonsByModuleIdRequest{
		ModuleId: moduleId,
	})
}

func (c *CourseClient) UpdateLesson(ctx context.Context, id, title, description string, position int64) (*lessonservicepb.UpdateLessonResponse, error) {
	return c.lesson.UpdateLesson(ctx, &lessonservicepb.UpdateLessonRequest{
		Id:          id,
		Title:       title,
		Description: description,
		Position:    position,
	})
}

func (c *CourseClient) DeleteLesson(ctx context.Context, id string) (*lessonservicepb.DeleteLessonResponse, error) {
	return c.lesson.DeleteLesson(ctx, &lessonservicepb.DeleteLessonRequest{
		Id: id,
	})
}
