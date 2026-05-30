package main

import (
	"log"

	"github.com/TwiLightDM/diploma-gateway/internal/app"
	"github.com/TwiLightDM/diploma-gateway/internal/config"
)

// @title Course Service API
// @version 1.0
// @description API для пользователей Course Service
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.Load()

	if err := app.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
