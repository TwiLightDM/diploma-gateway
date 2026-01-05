package main

import (
	"github.com/TwiLightDM/diploma-gateway/internal/app"
	"github.com/TwiLightDM/diploma-gateway/internal/config"
	"log"
)

func main() {
	cfg := config.Load()

	if err := app.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
