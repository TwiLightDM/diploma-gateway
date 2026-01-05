package app

import (
	"context"
	"errors"
	"github.com/TwiLightDM/diploma-gateway/internal/config"
	"github.com/TwiLightDM/diploma-gateway/internal/grpc"
	"github.com/TwiLightDM/diploma-gateway/internal/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run(cfg *config.Config) error {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	userClient := grpc.NewUserClient(cfg.UserGRPCAddr)
	userHandler := handlers.NewUserHandler(userClient)

	defer func() {
		log.Println("Closing gRPC connection to user service")
		_ = userClient.Close()
	}()

	registerRoutes(e, userHandler)

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

func registerRoutes(e *echo.Echo, userHandler *handlers.UserHandler) {
	public := e.Group("/auth")
	public.POST("/login", userHandler.Login)
	public.POST("/signup", userHandler.SignUp)

	users := e.Group("/users")
	users.GET("/:id", userHandler.ReadUser)
	users.PATCH("/:id", userHandler.UpdateUser)
	users.PATCH("/:id/password", userHandler.ChangePassword)
}
