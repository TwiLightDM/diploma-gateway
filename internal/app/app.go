package app

import (
	"context"
	"errors"
	"github.com/TwiLightDM/diploma-gateway/internal/config"
	"github.com/TwiLightDM/diploma-gateway/internal/grpc/course-service"
	"github.com/TwiLightDM/diploma-gateway/internal/grpc/user-service"
	"github.com/TwiLightDM/diploma-gateway/internal/handlers"
	"github.com/TwiLightDM/diploma-gateway/internal/middlewares"
	"github.com/TwiLightDM/diploma-gateway/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) error {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	jwtService := services.NewJWTService(cfg.JWTSecret)

	authMiddleware := middlewares.AuthMiddleware(jwtService)

	userClient := user_service.NewUserClient(cfg.UserGRPCAddr)
	userHandler := handlers.NewUserHandler(userClient)

	courseClient := course_service.NewCourseClient(cfg.CourseGRPCAddr)
	courseHandler := handlers.NewCourseHandler(courseClient)
	moduleHandler := handlers.NewModuleHandler(courseClient)
	lessonHandler := handlers.NewLessonHandler(courseClient)

	defer func() {
		log.Println("Closing gRPC connection to user service")
		_ = userClient.Close()
	}()

	registerRoutes(e, authMiddleware, userHandler, courseHandler, moduleHandler, lessonHandler)

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
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down gateway...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		return err
	}

	log.Println("Gateway stopped gracefully")
	return nil
}

func registerRoutes(e *echo.Echo,
	authMiddleware echo.MiddlewareFunc,
	userHandler *handlers.UserHandler,
	courseHandler *handlers.CourseHandler,
	moduleHandler *handlers.ModuleHandler,
	lessonHandler *handlers.LessonHandler,
) {
	public := e.Group("/auth")
	public.POST("/login", userHandler.Login)
	public.POST("/signup", userHandler.SignUp)

	users := e.Group("/users", authMiddleware)
	users.GET("", userHandler.ReadUser)
	users.GET("/:id", userHandler.ReadSelf)
	users.PATCH("/:id", userHandler.UpdateUser)
	users.PATCH("/:id/password", userHandler.ChangePassword)
	users.GET("/:id/courses", courseHandler.ReadAllCoursesByOwnerId)

	courses := e.Group("/courses", authMiddleware)
	courses.POST("", courseHandler.CreateCourse)
	courses.GET("/:id", courseHandler.ReadCourse)
	courses.PATCH("/:id", courseHandler.UpdateCourse)
	courses.PATCH("/:id/publish", courseHandler.UpdatePublishedAt)
	courses.DELETE("/:id", courseHandler.DeleteCourse)
	courses.GET("/:id/modules", moduleHandler.ReadAllModulesByCourseId)

	modules := e.Group("/modules", authMiddleware)
	modules.POST("", moduleHandler.CreateModule)
	modules.GET("/:id", moduleHandler.ReadModule)
	modules.PATCH("/:id", moduleHandler.UpdateModule)
	modules.DELETE("/:id", moduleHandler.DeleteModule)
	modules.GET("/:id/lessons", lessonHandler.ReadAllLessonsByCourseId)

	lessons := e.Group("/lessons", authMiddleware)
	lessons.POST("", lessonHandler.CreateLesson)
	lessons.GET("/:id", lessonHandler.ReadLesson)
	lessons.PATCH("/:id", lessonHandler.UpdateLesson)
	lessons.DELETE("/:id", lessonHandler.DeleteLesson)
}
