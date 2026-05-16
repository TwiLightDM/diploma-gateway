package course_service

import (
	"context"

	"github.com/TwiLightDM/diploma-course-service/proto/lessonprogressservicepb"
)

func (c *CourseClient) CreateLessonProgress(ctx context.Context, userId, lessonId string) (*lessonprogressservicepb.CreateLessonProgressResponse, error) {
	return c.lessonProgress.CreateLessonProgress(ctx, &lessonprogressservicepb.CreateLessonProgressRequest{
		UserId:   userId,
		LessonId: lessonId,
	})
}

func (c *CourseClient) ReadLessonProgressByUserId(ctx context.Context, userId string) (*lessonprogressservicepb.ReadLessonProgressByUserIdResponse, error) {
	return c.lessonProgress.ReadLessonProgressByUserId(ctx, &lessonprogressservicepb.ReadLessonProgressByUserIdRequest{
		UserId: userId,
	})
}

func (c *CourseClient) ReadLessonProgressByUserIdAndLessonId(ctx context.Context, userId, lessonId string) (*lessonprogressservicepb.ReadLessonProgressByUserIdAndLessonIdResponse, error) {
	return c.lessonProgress.ReadLessonProgressByUserIdAndLessonId(ctx, &lessonprogressservicepb.ReadLessonProgressByUserIdAndLessonIdRequest{
		UserId:   userId,
		LessonId: lessonId,
	})
}

func (c *CourseClient) ReadCourseProgressByUserId(ctx context.Context, userId, courseId string) (*lessonprogressservicepb.ReadCourseProgressByUserIdResponse, error) {
	return c.lessonProgress.ReadCourseProgressByUserId(ctx, &lessonprogressservicepb.ReadCourseProgressByUserIdRequest{
		UserId:   userId,
		CourseId: courseId,
	})
}

func (c *CourseClient) ReadCourseStatistics(ctx context.Context, courseId string) (*lessonprogressservicepb.ReadCourseStatisticsResponse, error) {
	return c.lessonProgress.ReadCourseStatistics(ctx, &lessonprogressservicepb.ReadCourseStatisticsRequest{
		CourseId: courseId,
	})
}

func (c *CourseClient) ReadModuleProgressByUserId(ctx context.Context, userId, moduleId string) (*lessonprogressservicepb.ReadModuleProgressByUserIdResponse, error) {
	return c.lessonProgress.ReadModuleProgressByUserId(ctx, &lessonprogressservicepb.ReadModuleProgressByUserIdRequest{
		UserId:   userId,
		ModuleId: moduleId,
	})
}

func (c *CourseClient) ReadModuleStatistics(ctx context.Context, moduleId string) (*lessonprogressservicepb.ReadModuleStatisticsResponse, error) {
	return c.lessonProgress.ReadModuleStatistics(ctx, &lessonprogressservicepb.ReadModuleStatisticsRequest{
		ModuleId: moduleId,
	})
}
