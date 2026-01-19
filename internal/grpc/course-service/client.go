package course_service

import (
	"github.com/TwiLightDM/diploma-course-service/proto/courseservicepb"
	"github.com/TwiLightDM/diploma-course-service/proto/lessonservicepb"
	"github.com/TwiLightDM/diploma-course-service/proto/moduleservicepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type CourseClient struct {
	course courseservicepb.CourseServiceClient
	module moduleservicepb.ModuleServiceClient
	lesson lessonservicepb.LessonServiceClient
	conn   *grpc.ClientConn
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
		course: courseservicepb.NewCourseServiceClient(conn),
		module: moduleservicepb.NewModuleServiceClient(conn),
		lesson: lessonservicepb.NewLessonServiceClient(conn),
		conn:   conn,
	}
}

func (c *CourseClient) Close() error {
	return c.conn.Close()
}
