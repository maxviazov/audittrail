package main

import (
	"fmt"
	"github.com/maxviazov/audittrail/internal/config"
	"github.com/maxviazov/audittrail/pkg/logger"
	"log"
)

func main() {
	fmt.Println("==> Starting server <==")

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	appLogger, err := logger.New(cfg.Log.ConsoleLevel, cfg.Log.FileLevel)
	if err != nil {
		log.Fatalf("error creating logger: %v", err)
	}
	appLogger.Info().Msg("Server started successfully")

	appLogger.Info().Msgf("Console log level: %s", cfg.Log.ConsoleLevel)
	appLogger.Info().Msgf("File log level: %s", cfg.Log.FileLevel)

}
