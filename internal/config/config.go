package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	GatewayPort    string
	UserGRPCAddr   string
	CourseGRPCAddr string
	JWTSecret      string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file didn't found")
	}

	cfg := &Config{}

	cfg.GatewayPort = os.Getenv("GATEWAY_PORT")
	cfg.UserGRPCAddr = os.Getenv("USER_GRPC_ADDR")
	cfg.CourseGRPCAddr = os.Getenv("COURSE_GRPC_ADDR")
	cfg.JWTSecret = os.Getenv("JWT_SECRET")

	return cfg
}
