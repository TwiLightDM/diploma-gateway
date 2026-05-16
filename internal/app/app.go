package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/TwiLightDM/diploma-gateway/internal/config"
	"github.com/TwiLightDM/diploma-gateway/internal/grpc/course-service"
	"github.com/TwiLightDM/diploma-gateway/internal/grpc/user-service"
	"github.com/TwiLightDM/diploma-gateway/internal/handlers"
	"github.com/TwiLightDM/diploma-gateway/internal/middlewares"
	"github.com/TwiLightDM/diploma-gateway/internal/services"
	"github.com/TwiLightDM/diploma-gateway/package/databases/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(cfg *config.Config) error {
	err := postgres.RunMigrations(cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Database)
	if err != nil {
		return err
	}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestLogger())

	jwtService := services.NewJWTService(cfg.JWTSecret)

	authMiddleware := middlewares.AuthMiddleware(jwtService)

	userClient := user_service.NewUserClient(cfg.UserGRPCAddr)
	userHandler := handlers.NewUserHandler(userClient)
	groupHandler := handlers.NewGroupHandler(userClient)
	groupMemberHandler := handlers.NewGroupMemberHandler(userClient)

	courseClient := course_service.NewCourseClient(cfg.CourseGRPCAddr)
	courseHandler := handlers.NewCourseHandler(courseClient)
	moduleHandler := handlers.NewModuleHandler(courseClient)
	lessonHandler := handlers.NewLessonHandler(courseClient)
	lessonFileHandler := handlers.NewLessonFileHandler(courseClient)
	groupCourseHandler := handlers.NewGroupCourseHandler(courseClient)
	taskHandler := handlers.NewTaskHandler(courseClient)
	taskAttemptHandler := handlers.NewTaskAttemptHandler(courseClient)
	lessonProgressHandler := handlers.NewLessonProgressHandler(courseClient)
	completedCourseHandler := handlers.NewCompletedCourseHandler(courseClient)
	completedModuleHandler := handlers.NewCompletedModuleHandler(courseClient)
	completedTheoryCourseHandler := handlers.NewCompletedTheoryCourseHandler(courseClient)

	defer func() {
		log.Println("Closing gRPC connection to user service")
		_ = userClient.Close()
	}()

	registerRoutes(
		e,
		authMiddleware,
		userHandler,
		groupHandler,
		groupMemberHandler,
		courseHandler,
		moduleHandler,
		lessonHandler,
		lessonFileHandler,
		groupCourseHandler,
		taskHandler,
		taskAttemptHandler,
		lessonProgressHandler,
		completedCourseHandler,
		completedModuleHandler,
		completedTheoryCourseHandler,
	)

	server := &http.Server{
		Addr:    ":" + cfg.GatewayPort,
		Handler: e,
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	go func() {
		log.Printf("Gateway started on :%s", cfg.GatewayPort)
		if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down gateway...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	log.Println("Gateway stopped gracefully")
	return nil
}

func registerRoutes(e *echo.Echo,
	authMiddleware echo.MiddlewareFunc,
	userHandler *handlers.UserHandler,
	groupHandler *handlers.GroupHandler,
	groupMemberHandler *handlers.GroupMemberHandler,
	courseHandler *handlers.CourseHandler,
	moduleHandler *handlers.ModuleHandler,
	lessonHandler *handlers.LessonHandler,
	lessonFileHandler *handlers.LessonFileHandler,
	groupCourseHandler *handlers.GroupCourseHandler,
	taskHandler *handlers.TaskHandler,
	taskAttemptHandler *handlers.TaskAttemptHandler,
	lessonProgressHandler *handlers.LessonProgressHandler,
	completedCourseHandler *handlers.CompletedCourseHandler,
	completedModuleHandler *handlers.CompletedModuleHandler,
	completedTheoryCourseHandler *handlers.CompletedTheoryCourseHandler,
) {
	public := e.Group("/auth")
	public.POST("/login", userHandler.Login)
	public.POST("/signup", userHandler.SignUp)
	public.POST("/refresh", userHandler.Refresh, authMiddleware)

	users := e.Group("/users", authMiddleware)
	users.GET("/me", userHandler.ReadSelf)
	users.GET("/all", userHandler.ReadAllUser)
	users.GET("/:id", userHandler.ReadUser)
	users.PATCH("", userHandler.UpdateUser)
	users.PATCH("/role", userHandler.UpdateUserRole)
	users.PATCH("/password", userHandler.ChangePassword)

	groups := e.Group("/groups", authMiddleware)
	groups.POST("", groupHandler.CreateGroup)
	groups.GET("/my", groupHandler.ReadAllGroupsByOwnerId)
	groups.GET("/:id", groupHandler.ReadGroup)
	groups.PATCH("/:id", groupHandler.UpdateGroup)
	groups.DELETE("/:id", groupHandler.DeleteGroup)

	groupMembers := e.Group("/group-members", authMiddleware)
	groupMembers.POST("", groupMemberHandler.CreateGroupMember)
	groupMembers.GET("", groupMemberHandler.ReadAllGroupMembersByUserId)
	groupMembers.GET("/:id", groupMemberHandler.ReadAllGroupMembersByGroupId)
	groupMembers.DELETE("/:id", groupMemberHandler.DeleteGroupMember)

	groupCourses := e.Group("/group-courses", authMiddleware)
	groupCourses.POST("", groupCourseHandler.CreateGroupCourse)
	groupCourses.GET("/:id", groupCourseHandler.ReadAllGroupCoursesByCourseId)
	groupCourses.DELETE("/:id", groupCourseHandler.DeleteGroupCourse)

	courses := e.Group("/courses", authMiddleware)
	courses.POST("", courseHandler.CreateCourse)
	courses.GET("", courseHandler.ReadAllCourses)
	courses.GET("/my", courseHandler.ReadAllCoursesByOwnerId)
	courses.GET("/:id", courseHandler.ReadCourse)
	courses.PATCH("/:id", courseHandler.UpdateCourse)
	courses.PATCH("/:id/publish", courseHandler.UpdatePublishedAt)
	courses.DELETE("/:id", courseHandler.DeleteCourse)

	modules := e.Group("/modules", authMiddleware)
	modules.POST("", moduleHandler.CreateModule)
	modules.GET("/courses/:course_id", moduleHandler.ReadAllModulesByCourseId)
	modules.GET("/:id", moduleHandler.ReadModule)
	modules.PATCH("/:id", moduleHandler.UpdateModule)
	modules.DELETE("/:id", moduleHandler.DeleteModule)

	lessons := e.Group("/lessons", authMiddleware)
	lessons.POST("", lessonHandler.CreateLesson)
	lessons.GET("/modules/:module_id", lessonHandler.ReadAllLessonsByCourseId)
	lessons.GET("/:id", lessonHandler.ReadLesson)
	lessons.PATCH("/:id", lessonHandler.UpdateLesson)
	lessons.DELETE("/:id", lessonHandler.DeleteLesson)

	lessonFiles := e.Group("/lessons/files", authMiddleware)
	lessonFiles.POST("", lessonFileHandler.UploadFile)
	lessonFiles.GET("/:lesson_id", lessonFileHandler.GetFiles)
	lessonFiles.DELETE("/:id", lessonFileHandler.DeleteFile)

	tasks := e.Group("/tasks", authMiddleware)
	tasks.POST("", taskHandler.CreateTask)
	tasks.GET("/modules/:module_id", taskHandler.ReadTasksByModuleId)
	tasks.GET("/courses/:course_id", taskHandler.ReadTasksByCourseId)
	tasks.GET("/:id", taskHandler.ReadTask)
	tasks.PATCH("/:id", taskHandler.UpdateTask)
	tasks.DELETE("/:id", taskHandler.DeleteTask)

	taskAttempts := e.Group("/task-attempts", authMiddleware)
	taskAttempts.POST("", taskAttemptHandler.SubmitTaskAttempt)
	taskAttempts.GET("/courses/:course_id/my", taskAttemptHandler.ReadAllTaskAttemptsByYourIdAndCourseId)
	taskAttempts.GET("/modules/:module_id/my", taskAttemptHandler.ReadAllTaskAttemptsByYourIdAndModuleId)
	taskAttempts.GET("/courses/:course_id", taskAttemptHandler.ReadAllTaskAttemptsByCourseId)
	taskAttempts.GET("/modules/:module_id", taskAttemptHandler.ReadAllTaskAttemptsByModuleId)
	taskAttempts.GET("/users/:user_id/courses/:course_id", taskAttemptHandler.ReadAllTaskAttemptsByUserIdAndCourseId)
	taskAttempts.GET("/users/:user_id/modules/:module_id", taskAttemptHandler.ReadAllTaskAttemptsByUserIdAndModuleId)
	taskAttempts.GET("/:id", taskAttemptHandler.ReadTaskAttempt)

	lessonProgress := e.Group("/lesson-progresses", authMiddleware)
	lessonProgress.POST("", lessonProgressHandler.CreateLessonProgress)
	lessonProgress.GET("", lessonProgressHandler.ReadLessonProgressByUserId)
	lessonProgress.GET("/check", lessonProgressHandler.ReadLessonProgressByUserIdAndLessonId)

	courseProgress := e.Group("/progresses", authMiddleware)
	courseProgress.GET("/courses/:courseId", lessonProgressHandler.ReadCourseProgress)
	courseProgress.GET("/courses/:courseId/statistics", lessonProgressHandler.ReadCourseStatistics)

	moduleProgress := e.Group("/progresses", authMiddleware)
	moduleProgress.GET("/modules/:moduleId", lessonProgressHandler.ReadModuleProgress)
	moduleProgress.GET("/modules/:moduleId/statistics", lessonProgressHandler.ReadModuleStatistics)

	completedCourse := e.Group("/completed-courses", authMiddleware)
	completedCourse.POST("", completedCourseHandler.CreateCompletedCourse)
	completedCourse.GET("/my", completedCourseHandler.ReadAllCompletedCoursesByUserId)
	completedCourse.GET("/courses/:courseId", completedCourseHandler.ReadAllCompletedCoursesByCourseId)
	completedCourse.GET("/courses/:courseId/my", completedCourseHandler.ReadCompletedCourseByUserIdAndCourseId)

	completedModule := e.Group("/completed-modules", authMiddleware)
	completedModule.POST("", completedCourseHandler.CreateCompletedCourse)
	completedModule.GET("/my", completedModuleHandler.ReadAllCompletedModulesByUserId)
	completedModule.GET("/modules/:moduleId", completedModuleHandler.ReadAllCompletedModulesByModuleId)
	completedModule.GET("/modules/:moduleId/my", completedModuleHandler.ReadCompletedModuleByUserIdAndModuleId)

	completedTheoryCourse := e.Group("/completed-theory-courses", authMiddleware)
	completedTheoryCourse.POST("", completedTheoryCourseHandler.CreateCompletedTheoryCourse)
	completedTheoryCourse.GET("/my", completedTheoryCourseHandler.ReadAllCompletedTheoryCoursesByUserId)
	completedTheoryCourse.GET("/courses/:courseId", completedTheoryCourseHandler.ReadAllCompletedTheoryCoursesByCourseId)
	completedTheoryCourse.GET("/courses/:courseId/my", completedTheoryCourseHandler.ReadCompletedTheoryCourseByUserIdAndCourseId)
}
