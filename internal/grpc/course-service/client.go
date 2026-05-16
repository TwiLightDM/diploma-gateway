package course_service

import (
	"log"

	"github.com/TwiLightDM/diploma-course-service/proto/completedcourseservicepb"
	"github.com/TwiLightDM/diploma-course-service/proto/completedmoduleservicepb"
	"github.com/TwiLightDM/diploma-course-service/proto/completedtheorycourseservicepb"
	"github.com/TwiLightDM/diploma-course-service/proto/courseservicepb"
	"github.com/TwiLightDM/diploma-course-service/proto/groupcourseservicepb"
	"github.com/TwiLightDM/diploma-course-service/proto/lessonfileservicepb"
	"github.com/TwiLightDM/diploma-course-service/proto/lessonprogressservicepb"
	"github.com/TwiLightDM/diploma-course-service/proto/lessonservicepb"
	"github.com/TwiLightDM/diploma-course-service/proto/moduleservicepb"
	"github.com/TwiLightDM/diploma-course-service/proto/taskattemptservicepb"
	"github.com/TwiLightDM/diploma-course-service/proto/taskservicepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CourseClient struct {
	course                courseservicepb.CourseServiceClient
	module                moduleservicepb.ModuleServiceClient
	lesson                lessonservicepb.LessonServiceClient
	lessonFile            lessonfileservicepb.LessonFileServiceClient
	groupCourse           groupcourseservicepb.GroupCourseServiceClient
	task                  taskservicepb.TaskServiceClient
	taskAttempt           taskattemptservicepb.TaskAttemptServiceClient
	lessonProgress        lessonprogressservicepb.LessonProgressServiceClient
	completedCourse       completedcourseservicepb.CompletedCourseServiceClient
	completedModule       completedmoduleservicepb.CompletedModuleServiceClient
	completedTheoryCourse completedtheorycourseservicepb.CompletedTheoryCourseServiceClient
	conn                  *grpc.ClientConn
}

func NewCourseClient(address string) *CourseClient {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("failed to connect to course service: %v", err)
	}

	return &CourseClient{
		course:                courseservicepb.NewCourseServiceClient(conn),
		module:                moduleservicepb.NewModuleServiceClient(conn),
		lesson:                lessonservicepb.NewLessonServiceClient(conn),
		lessonFile:            lessonfileservicepb.NewLessonFileServiceClient(conn),
		groupCourse:           groupcourseservicepb.NewGroupCourseServiceClient(conn),
		task:                  taskservicepb.NewTaskServiceClient(conn),
		taskAttempt:           taskattemptservicepb.NewTaskAttemptServiceClient(conn),
		lessonProgress:        lessonprogressservicepb.NewLessonProgressServiceClient(conn),
		completedCourse:       completedcourseservicepb.NewCompletedCourseServiceClient(conn),
		completedModule:       completedmoduleservicepb.NewCompletedModuleServiceClient(conn),
		completedTheoryCourse: completedtheorycourseservicepb.NewCompletedTheoryCourseServiceClient(conn),
		conn:                  conn,
	}
}

func (c *CourseClient) Close() error {
	return c.conn.Close()
}
